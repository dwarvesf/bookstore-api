package book

import (
	"context"
	"errors"
	"net/http/httptest"
	"testing"

	bookMocks "github.com/dwarvesf/bookstore-api/mocks/pkg/repository/book"
	topicMocks "github.com/dwarvesf/bookstore-api/mocks/pkg/repository/topic"
	"github.com/dwarvesf/bookstore-api/pkg/config"
	"github.com/dwarvesf/bookstore-api/pkg/handler/testutil"
	"github.com/dwarvesf/bookstore-api/pkg/logger/monitor"
	"github.com/dwarvesf/bookstore-api/pkg/middleware"
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/repository"
	"github.com/dwarvesf/bookstore-api/pkg/repository/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_impl_UpdateBook(t *testing.T) {
	name := "book1"
	author := "author1"
	topicID := 1

	type mocked struct {
		expGetByUserAndIDCalled bool
		getByUserAndID          *model.Book
		getByUserAndIDErr       error
		expIsTopicExistCalled   bool
		isTopicExist            bool
		isTopicExistErr         error
		expUpdateBookCalled     bool
		updateBook              *model.Book
		updateBooksErr          error
		expGetBookCalled        bool
		getBook                 *model.Book
		getBookErr              error
	}
	type args struct {
		uID int
		req model.UpdateBookRequest
	}
	tests := map[string]struct {
		mocked  mocked
		args    args
		wantErr bool
	}{
		"success": {
			mocked: mocked{
				expGetByUserAndIDCalled: true,
				getByUserAndID: &model.Book{
					ID:     1,
					Name:   "book1",
					Author: "author1",
				},
				expIsTopicExistCalled: true,
				isTopicExist:          true,
				expUpdateBookCalled:   true,
				updateBook: &model.Book{
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
				req: model.UpdateBookRequest{
					ID:      1,
					Name:    &name,
					Author:  &author,
					TopicID: &topicID,
				},
			},
			wantErr: false,
		},
		"topic not found": {
			mocked: mocked{
				expGetByUserAndIDCalled: true,
				getByUserAndID: &model.Book{
					ID:     1,
					Name:   "book1",
					Author: "author1",
				},
				expIsTopicExistCalled: true,
				isTopicExist:          false,
				expUpdateBookCalled:   false,
			},
			args: args{
				uID: 1,
				req: model.UpdateBookRequest{
					ID:      1,
					Name:    &name,
					Author:  &author,
					TopicID: &topicID,
				},
			},
			wantErr: true,
		},
		"update book failed": {
			mocked: mocked{
				expGetByUserAndIDCalled: true,
				getByUserAndID: &model.Book{
					ID:     1,
					Name:   "book1",
					Author: "author1",
				},
				expIsTopicExistCalled: true,
				isTopicExist:          true,
				expUpdateBookCalled:   true,
				updateBooksErr:        errors.New("error"),
			},
			args: args{
				uID: 1,
				req: model.UpdateBookRequest{
					ID:      1,
					Name:    &name,
					Author:  &author,
					TopicID: &topicID,
				},
			},
			wantErr: true,
		},
		"not found book": {
			mocked: mocked{
				expGetByUserAndIDCalled: true,
				getByUserAndIDErr:       errors.New("error"),
			},
			args: args{
				uID: 1,
				req: model.UpdateBookRequest{
					ID:      1,
					Name:    &name,
					Author:  &author,
					TopicID: &topicID,
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

			if tt.mocked.expGetByUserAndIDCalled {
				bookRepoMock.
					EXPECT().
					GetByUserAndID(mock.Anything, mock.Anything, mock.Anything).
					Return(tt.mocked.getByUserAndID, tt.mocked.getByUserAndIDErr)
			}

			if tt.mocked.expIsTopicExistCalled {
				topicRepoMock.
					EXPECT().
					IsExist(mock.Anything, mock.Anything).
					Return(tt.mocked.isTopicExist, tt.mocked.isTopicExistErr)
			}

			if tt.mocked.expUpdateBookCalled {
				bookRepoMock.
					EXPECT().
					Update(mock.Anything, mock.Anything).
					Return(tt.mocked.updateBook, tt.mocked.updateBooksErr)
			}

			if tt.mocked.expGetBookCalled {
				bookRepoMock.
					EXPECT().
					GetByID(mock.Anything, mock.Anything).
					Return(tt.mocked.getBook, tt.mocked.getBookErr)
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
			ctx := context.WithValue(ginCtx.Request.Context(), middleware.UserIDCtxKey, tt.args.uID)
			ginCtx.Request = ginCtx.Request.WithContext(ctx)
			books, err := c.UpdateBook(ginCtx.Request.Context(), tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.Updatebook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.Equal(t, tt.mocked.getBook, books)
			}
		})
	}
}
