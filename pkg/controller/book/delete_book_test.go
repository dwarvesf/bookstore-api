package book

import (
	"context"
	"errors"
	"net/http/httptest"
	"testing"

	bookMocks "github.com/dwarvesf/bookstore-api/mocks/pkg/repository/book"
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

func Test_impl_DeleteBook(t *testing.T) {
	type mocked struct {
		expIsBookExistCalled bool
		isBookExist          bool
		isBookExistErr       error
		expDeleteBookCalled  bool
		deleteBooksErr       error
	}
	type args struct {
		uID int
		id  int
	}
	tests := map[string]struct {
		mocked  mocked
		args    args
		err     error
		wantErr bool
	}{
		"success": {
			mocked: mocked{
				expIsBookExistCalled: true,
				isBookExist:          true,
				expDeleteBookCalled:  true,
			},
			args: args{
				uID: 1,
				id:  1,
			},
			wantErr: false,
		},
		"book not found": {
			mocked: mocked{
				expIsBookExistCalled: true,
				isBookExist:          false,
			},
			args: args{
				uID: 1,
				id:  1,
			},
			err:     model.ErrBookNotFound,
			wantErr: true,
		},
		"failed to delete": {
			mocked: mocked{
				expIsBookExistCalled: true,
				isBookExist:          true,
				expDeleteBookCalled:  true,
				deleteBooksErr:       errors.New("error"),
			},
			args: args{
				uID: 1,
				id:  1,
			},
			err:     errors.New("error"),
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var (
				bookRepoMock = bookMocks.NewRepo(t)
			)

			if tt.mocked.expIsBookExistCalled {
				bookRepoMock.
					EXPECT().
					IsExist(mock.Anything, mock.Anything, mock.Anything).
					Return(tt.mocked.isBookExist, tt.mocked.isBookExistErr)
			}

			if tt.mocked.expDeleteBookCalled {
				bookRepoMock.
					EXPECT().
					Delete(mock.Anything, mock.Anything).
					Return(tt.mocked.deleteBooksErr)
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

			w := httptest.NewRecorder()
			ginCtx := testutil.NewRequest(w, testutil.MethodDelete, nil, nil, nil, nil)
			ctx := context.WithValue(ginCtx.Request.Context(), middleware.UserIDCtxKey, tt.args.id)
			ginCtx.Request = ginCtx.Request.WithContext(ctx)
			err = c.DeleteBook(ginCtx.Request.Context(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.DeleteBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.err, err)
		})
	}
}
