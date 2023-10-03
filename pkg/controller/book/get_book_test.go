package book

import (
	"context"
	"errors"
	"testing"

	mocks "github.com/dwarvesf/bookstore-api/mocks/pkg/repository/book"
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

func Test_impl_GetBook(t *testing.T) {
	type mocked struct {
		expGetBookCalled bool
		getBook          *model.Book
		getBookErr       error
	}
	type args struct {
		uID    int
		bookID int
	}
	tests := map[string]struct {
		mocked  mocked
		args    args
		wantErr bool
	}{
		"success": {
			mocked: mocked{
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
				uID:    1,
				bookID: 1,
			},
			wantErr: false,
		},
		"failed": {
			mocked: mocked{
				expGetBookCalled: true,
				getBookErr:       errors.New("error"),
			},
			args: args{
				uID:    1,
				bookID: 1,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var (
				bookRepoMock = mocks.NewRepo(t)
			)

			if tt.mocked.expGetBookCalled {
				bookRepoMock.
					EXPECT().
					GetByUserAndID(mock.Anything, mock.Anything, mock.Anything).
					Return(tt.mocked.getBook, tt.mocked.getBookErr)
			}

			c := &impl{
				repo: &repository.Repo{
					Book: bookRepoMock,
				},
				cfg:     config.LoadTestConfig(),
				monitor: monitor.TestMonitor(),
			}

			_, err := db.Init(c.cfg)
			require.NoError(t, err)

			ctx := context.Background()
			ctx = context.WithValue(ctx, middleware.UserIDCtxKey, tt.args.uID)

			book, err := c.GetBook(ctx, tt.args.bookID)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.GetBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.mocked.getBook, book)
		})
	}
}
