package portal

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"net/url"
	"testing"

	mocks "github.com/dwarvesf/bookstore-api/mocks/pkg/controller/book"
	"github.com/dwarvesf/bookstore-api/pkg/config"
	"github.com/dwarvesf/bookstore-api/pkg/handler/testutil"
	"github.com/dwarvesf/bookstore-api/pkg/handler/v1/view"
	"github.com/dwarvesf/bookstore-api/pkg/logger"
	"github.com/dwarvesf/bookstore-api/pkg/logger/monitor"
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_GetBook(t *testing.T) {
	type mocked struct {
		page        string
		pageSize    string
		sort        string
		query       string
		expGetBooks bool
		books       *model.ListResult[model.Book]
		getBooksErr error
	}

	type expected struct {
		Status  int
		Body    view.GetBooksResponse
		WantErr bool
		Err     string
	}
	tests := map[string]struct {
		mocked   mocked
		expected expected
	}{
		"success": {
			mocked: mocked{
				page:        "1",
				pageSize:    "10",
				sort:        "created_at desc",
				query:       "",
				expGetBooks: true,
				books: &model.ListResult[model.Book]{
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
			expected: expected{
				Status: 200,
				Body: view.GetBooksResponse{
					Metadata: view.Metadata{
						Page:       1,
						PageSize:   10,
						TotalPages: 2,
						Sort:       "created_at desc",
						HasNext:    false,
					},
					Data: []*view.Book{
						{
							ID:     1,
							Name:   "book1",
							Author: "author1",
							Topic: &view.Topic{
								ID:   1,
								Name: "topic1",
								Code: "code1",
							},
						},
						{
							ID:     2,
							Name:   "book2",
							Author: "author2",
							Topic: &view.Topic{
								ID:   2,
								Name: "topic2",
								Code: "code2",
							},
						},
					},
				},
			},
		},
		"failed to get": {
			mocked: mocked{
				page:        "1",
				pageSize:    "10",
				sort:        "created_at desc",
				query:       "",
				expGetBooks: true,
				getBooksErr: errors.New("failed to get books"),
			},
			expected: expected{
				Status:  500,
				WantErr: true,
				Err:     "INTERNAL_ERROR",
			},
		},
	}
	for name, tt := range tests {
		w := httptest.NewRecorder()
		cfg := config.LoadTestConfig()
		query := url.Values{}
		if tt.mocked.page != "" {
			query.Add("page", tt.mocked.page)
		}
		if tt.mocked.pageSize != "" {
			query.Add("page_size", tt.mocked.pageSize)
		}
		if tt.mocked.sort != "" {
			query.Add("sort", tt.mocked.sort)
		}
		if tt.mocked.query != "" {
			query.Add("query", tt.mocked.query)
		}

		ginCtx := testutil.NewRequest(w, testutil.MethodGet, nil, nil, query, nil)

		var (
			ctrlMock = mocks.NewController(t)
		)

		if tt.mocked.expGetBooks {
			ctrlMock.EXPECT().GetBooks(mock.Anything, mock.Anything).Return(tt.mocked.books, tt.mocked.getBooksErr)
		}
		t.Run(name, func(t *testing.T) {
			h := Handler{
				log:      logger.NewLogger(),
				cfg:      cfg,
				bookCtrl: ctrlMock,
				monitor:  monitor.TestMonitor(),
			}
			h.GetBooks(ginCtx)

			assert.Equal(t, tt.expected.Status, w.Code)
			resBody := w.Body.String()
			body, err := json.Marshal(tt.expected.Body)
			assert.Nil(t, err)
			if !tt.expected.WantErr {
				assert.Equal(t, string(body), resBody)
			} else {
				assert.Contains(t, resBody, tt.expected.Err)
			}
		})
	}
}

func TestHandler_CreateBook(t *testing.T) {
	type mocked struct {
		expCreateBook bool
		books         *model.Book
		createBookErr error
	}

	type args struct {
		req view.CreateBookRequest
	}

	type expected struct {
		Status  int
		Body    *view.Book
		WantErr bool
		Err     string
	}
	tests := map[string]struct {
		mocked   mocked
		args     args
		expected expected
	}{
		"success": {
			mocked: mocked{
				expCreateBook: true,
				books: &model.Book{
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
				req: view.CreateBookRequest{
					Name:    "book1",
					Author:  "author1",
					TopicID: 1,
				},
			},
			expected: expected{
				Status: 200,
				Body: &view.Book{
					ID:     1,
					Name:   "book1",
					Author: "author1",
					Topic: &view.Topic{
						ID:   1,
						Name: "topic1",
						Code: "code1",
					},
				},
			},
		},
		"failed": {
			mocked: mocked{
				expCreateBook: true,
				createBookErr: errors.New("failed to create book"),
			},
			args: args{
				req: view.CreateBookRequest{
					Name:    "book1",
					Author:  "author1",
					TopicID: 1,
				},
			},
			expected: expected{
				Status:  500,
				WantErr: true,
				Err:     "INTERNAL_ERROR",
			},
		},
	}
	for name, tt := range tests {
		w := httptest.NewRecorder()
		cfg := config.LoadTestConfig()

		ginCtx := testutil.NewRequest(w, testutil.MethodGet, nil, nil, nil, tt.args.req)

		var (
			ctrlMock = mocks.NewController(t)
		)

		if tt.mocked.expCreateBook {
			ctrlMock.EXPECT().CreateBook(mock.Anything, mock.Anything).Return(tt.mocked.books, tt.mocked.createBookErr)
		}
		t.Run(name, func(t *testing.T) {
			h := Handler{
				log:      logger.NewLogger(),
				cfg:      cfg,
				bookCtrl: ctrlMock,
				monitor:  monitor.TestMonitor(),
			}
			h.CreateBook(ginCtx)

			assert.Equal(t, tt.expected.Status, w.Code)
			resBody := w.Body.String()
			body, err := json.Marshal(tt.expected.Body)
			assert.Nil(t, err)
			if !tt.expected.WantErr {
				assert.Equal(t, string(body), resBody)
			} else {
				assert.Contains(t, resBody, tt.expected.Err)
			}
		})
	}
}

func TestHandler_DeleteBook(t *testing.T) {
	type mocked struct {
		expDeleteBook bool
		deleteBookErr error
	}

	type args struct {
		id string
	}

	type expected struct {
		Status  int
		Body    *view.Message
		WantErr bool
		Err     string
	}
	tests := map[string]struct {
		mocked   mocked
		args     args
		expected expected
	}{
		"success": {
			mocked: mocked{
				expDeleteBook: true,
			},
			args: args{
				id: "1",
			},
			expected: expected{
				Status: 200,
				Body:   &view.Message{Message: "OK"},
			},
		},
		"failed": {
			mocked: mocked{
				expDeleteBook: true,
				deleteBookErr: errors.New("failed to create book"),
			},
			args: args{
				id: "1",
			},
			expected: expected{
				Status:  500,
				WantErr: true,
				Err:     "INTERNAL_ERROR",
			},
		},
	}
	for name, tt := range tests {
		w := httptest.NewRecorder()
		cfg := config.LoadTestConfig()
		param := []gin.Param{{Key: "id", Value: tt.args.id}}

		ginCtx := testutil.NewRequest(w, testutil.MethodGet, nil, param, nil, nil)

		var (
			ctrlMock = mocks.NewController(t)
		)

		if tt.mocked.expDeleteBook {
			ctrlMock.EXPECT().DeleteBook(mock.Anything, mock.Anything).Return(tt.mocked.deleteBookErr)
		}
		t.Run(name, func(t *testing.T) {
			h := Handler{
				log:      logger.NewLogger(),
				cfg:      cfg,
				bookCtrl: ctrlMock,
				monitor:  monitor.TestMonitor(),
			}
			h.DeleteBook(ginCtx)

			assert.Equal(t, tt.expected.Status, w.Code)
			resBody := w.Body.String()
			body, err := json.Marshal(tt.expected.Body)
			assert.Nil(t, err)
			if !tt.expected.WantErr {
				assert.Equal(t, string(body), resBody)
			} else {
				assert.Contains(t, resBody, tt.expected.Err)
			}
		})
	}
}

func TestHandler_UpdateBook(t *testing.T) {
	type mocked struct {
		expUpdateBook bool
		books         *model.Book
		updateBookErr error
	}

	type args struct {
		id  string
		req view.CreateBookRequest
	}

	type expected struct {
		Status  int
		Body    *view.Book
		WantErr bool
		Err     string
	}
	tests := map[string]struct {
		mocked   mocked
		args     args
		expected expected
	}{
		"success": {
			mocked: mocked{
				expUpdateBook: true,
				books: &model.Book{
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
				id: "1",
				req: view.CreateBookRequest{
					Name:    "book1",
					Author:  "author1",
					TopicID: 1,
				},
			},
			expected: expected{
				Status: 200,
				Body: &view.Book{
					ID:     1,
					Name:   "book1",
					Author: "author1",
					Topic: &view.Topic{
						ID:   1,
						Name: "topic1",
						Code: "code1",
					},
				},
			},
		},
		"failed": {
			mocked: mocked{
				expUpdateBook: true,
				updateBookErr: errors.New("failed to create book"),
			},
			args: args{
				id: "1",
				req: view.CreateBookRequest{
					Name:    "book1",
					Author:  "author1",
					TopicID: 1,
				},
			},
			expected: expected{
				Status:  500,
				WantErr: true,
				Err:     "INTERNAL_ERROR",
			},
		},
	}
	for name, tt := range tests {
		w := httptest.NewRecorder()
		cfg := config.LoadTestConfig()
		param := []gin.Param{{Key: "id", Value: tt.args.id}}

		ginCtx := testutil.NewRequest(w, testutil.MethodGet, nil, param, nil, tt.args.req)

		var (
			ctrlMock = mocks.NewController(t)
		)

		if tt.mocked.expUpdateBook {
			ctrlMock.EXPECT().UpdateBook(mock.Anything, mock.Anything).Return(tt.mocked.books, tt.mocked.updateBookErr)
		}
		t.Run(name, func(t *testing.T) {
			h := Handler{
				log:      logger.NewLogger(),
				cfg:      cfg,
				bookCtrl: ctrlMock,
				monitor:  monitor.TestMonitor(),
			}
			h.UpdateBook(ginCtx)

			assert.Equal(t, tt.expected.Status, w.Code)
			resBody := w.Body.String()
			body, err := json.Marshal(tt.expected.Body)
			assert.Nil(t, err)
			if !tt.expected.WantErr {
				assert.Equal(t, string(body), resBody)
			} else {
				assert.Contains(t, resBody, tt.expected.Err)
			}
		})
	}
}
