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

func Test_impl_GetBooks(t *testing.T) {
	type mocked struct {
		expGetBooksCalled bool
		getBooks          *model.ListResult[model.Book]
		getBooksErr       error
	}
	type args struct {
		uID     int
		topicID int
		req     model.ListQuery
	}
	tests := map[string]struct {
		mocked  mocked
		args    args
		wantErr bool
	}{
		"success": {
			mocked: mocked{
				expGetBooksCalled: true,
				getBooks: &model.ListResult[model.Book]{
					Pagination: model.Pagination{
						Page:       1,
						PageSize:   10,
						TotalPages: 2,
						Sort:       "created_at desc",
						HasNext:    false,
					},
					Data: []model.Book{
						{
							ID:     1,
							Name:   "book1",
							Author: "author1",
							Topic: &model.Topic{
								ID:   1,
								Name: "topic1",
								Code: "code1",
							},
						},
						{
							ID:     2,
							Name:   "book2",
							Author: "author2",
							Topic: &model.Topic{
								ID:   2,
								Name: "topic2",
								Code: "code2",
							},
						},
					},
				},
			},
			args: args{
				req: model.ListQuery{
					Page:     1,
					PageSize: 10,
					Sort:     "created_at desc",
				},
			},
			wantErr: false,
		},
		"failed": {
			mocked: mocked{
				expGetBooksCalled: true,
				getBooksErr:       errors.New("error"),
			},
			args: args{
				req: model.ListQuery{
					Page:     1,
					PageSize: 10,
					Sort:     "created_at desc",
				},
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var (
				bookRepoMock = mocks.NewRepo(t)
			)

			if tt.mocked.expGetBooksCalled {
				bookRepoMock.
					EXPECT().
					GetList(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(tt.mocked.getBooks, tt.mocked.getBooksErr)
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

			books, err := c.GetBooks(ctx, tt.args.req, tt.args.topicID)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.GetBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.mocked.getBooks, books)
		})
	}
}
