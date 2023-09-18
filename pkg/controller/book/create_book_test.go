package book

import (
	"context"
	"errors"
	"testing"

	bookMocks "github.com/dwarvesf/bookstore-api/mocks/pkg/repository/book"
	topicMocks "github.com/dwarvesf/bookstore-api/mocks/pkg/repository/topic"
	userMocks "github.com/dwarvesf/bookstore-api/mocks/pkg/repository/user"
	"github.com/dwarvesf/bookstore-api/pkg/config"
	"github.com/dwarvesf/bookstore-api/pkg/logger/monitor"
	"github.com/dwarvesf/bookstore-api/pkg/middleware"
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/repository"
	"github.com/dwarvesf/bookstore-api/pkg/repository/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_impl_CreateBooks(t *testing.T) {
	type mocked struct {
		expIsTopicExistCalled bool
		isTopicExist          bool
		isTopicExistErr       error
		expCreateBookCalled   bool
		createBook            *model.Book
		createBooksErr        error
		expGetBookCalled      bool
		getBook               *model.Book
		getBooksErr           error
	}
	type args struct {
		uID int
		req model.CreateBookRequest
	}
	tests := map[string]struct {
		mocked  mocked
		args    args
		wantErr bool
	}{
		"success": {
			mocked: mocked{
				expIsTopicExistCalled: true,
				isTopicExist:          true,
				expCreateBookCalled:   true,
				createBook: &model.Book{
					ID:     1,
					Name:   "book1",
					Author: "author1",
				},
				expGetBookCalled: true,
				getBook: &model.Book{
					ID:     1,
					Name:   "book1",
					Author: "author1",
					Topic: &model.Topic{
						ID:   1,
						Name: "topic1",
						Code: "code1",
					},
				},
			},
			args: args{
				uID: 1,
				req: model.CreateBookRequest{
					Name:    "book1",
					Author:  "author1",
					TopicID: 1,
				},
			},
			wantErr: false,
		},
		"topic not found": {
			mocked: mocked{
				expIsTopicExistCalled: true,
				isTopicExist:          false,
				expCreateBookCalled:   false,
				expGetBookCalled:      false,
			},
			args: args{
				uID: 1,
				req: model.CreateBookRequest{
					Name:    "book1",
					Author:  "author1",
					TopicID: 1,
				},
			},
			wantErr: true,
		},
		"create book failed": {
			mocked: mocked{
				expIsTopicExistCalled: true,
				isTopicExist:          true,
				expCreateBookCalled:   true,
				createBooksErr:        errors.New("error"),
				expGetBookCalled:      false,
			},
			args: args{
				uID: 1,
				req: model.CreateBookRequest{
					Name:    "book1",
					Author:  "author1",
					TopicID: 1,
				},
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var (
				bookRepoMock  = bookMocks.NewRepo(t)
				topicRepoMock = topicMocks.NewRepo(t)
				userRepoMock  = userMocks.NewRepo(t)
			)

			if tt.mocked.expIsTopicExistCalled {
				topicRepoMock.
					EXPECT().
					IsExist(mock.Anything, mock.Anything).
					Return(tt.mocked.isTopicExist, tt.mocked.isTopicExistErr)
			}

			if tt.mocked.expCreateBookCalled {
				bookRepoMock.
					EXPECT().
					Create(mock.Anything, mock.Anything).
					Return(tt.mocked.createBook, tt.mocked.createBooksErr)
			}

			if tt.mocked.expGetBookCalled {
				bookRepoMock.
					EXPECT().
					GetByID(mock.Anything, mock.Anything).
					Return(tt.mocked.getBook, tt.mocked.getBooksErr)
			}

			c := &impl{
				repo: &repository.Repo{
					Book:  bookRepoMock,
					Topic: topicRepoMock,
					User:  userRepoMock,
				},
				cfg:     config.LoadTestConfig(),
				monitor: monitor.TestMonitor(),
			}

			_, err := db.Init(c.cfg)
			require.NoError(t, err)

			ctx := context.Background()
			ctx = context.WithValue(ctx, middleware.UserIDCtxKey, tt.args.uID)

			books, err := c.CreateBook(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.CreateBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.Equal(t, tt.mocked.getBook.Topic, books.Topic)

				tt.mocked.getBook.Topic = nil
				books.Topic = nil
				assert.Equal(t, tt.mocked.getBook, books)
			}
		})
	}
}
