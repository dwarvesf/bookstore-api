package book

import (
	"errors"
	"net/http/httptest"
	"testing"

	bookMocks "github.com/dwarvesf/bookstore-api/mocks/pkg/repository/book"
	topicMocks "github.com/dwarvesf/bookstore-api/mocks/pkg/repository/topic"
	"github.com/dwarvesf/bookstore-api/pkg/config"
	"github.com/dwarvesf/bookstore-api/pkg/handler/testutil"
	"github.com/dwarvesf/bookstore-api/pkg/logger/monitor"
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
	}
	type args struct {
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
					Topic: &model.Topic{
						ID:   1,
						Name: "topic1",
						Code: "code1",
					},
				},
			},
			args: args{
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
			},
			args: args{
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
			},
			args: args{
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

			c := &impl{
				repo: &repository.Repo{
					Book:  bookRepoMock,
					Topic: topicRepoMock,
				},
				cfg:     config.LoadTestConfig(),
				monitor: monitor.TestMonitor(),
			}

			_, err := db.Init(c.cfg)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			ginCtx := testutil.NewRequest(w, testutil.MethodPost, nil, nil, nil, nil)
			books, err := c.CreateBook(ginCtx.Request.Context(), tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.CreateBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.mocked.createBook, books)
		})
	}
}
