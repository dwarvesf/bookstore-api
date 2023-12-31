package util

import (
	"database/sql"
	"net/http/httptest"
	"testing"

	"github.com/dwarvesf/bookstore-api/pkg/handler/testutil"
	"github.com/dwarvesf/bookstore-api/pkg/handler/v1/view"
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestHandleError(t *testing.T) {
	type args struct {
		err error
	}
	type expected struct {
		Status int
		Body   string
	}
	tests := map[string]struct {
		name     string
		args     args
		expected expected
	}{
		"valid error": {
			args: args{
				err: model.Error{Status: 400, Code: "bad_request", Message: "bad request"},
			},
			expected: expected{
				Status: 400,
				Body:   "bad request",
			},
		},
		"valid pointer error": {
			args: args{
				err: sql.ErrNoRows,
			},
			expected: expected{

				Status: 500,
				Body:   "no rows",
			},
		},
	}
	for name, tt := range tests {
		w := httptest.NewRecorder()
		ginCtx := testutil.NewRequest(w, testutil.MethodGet, nil, nil, nil, nil)
		t.Run(name, func(t *testing.T) {
			HandleError(ginCtx, tt.args.err)

			assert.Equal(t, tt.expected.Status, w.Code)
			resBody := w.Body.String()
			assert.Contains(t, resBody, tt.expected.Body)
		})
	}
}

func Test_tryParseError(t *testing.T) {
	type args struct {
		err error
	}
	tests := map[string]struct {
		name string
		args args
		want view.ErrorResponse
	}{
		"valid error": {
			args: args{
				err: model.Error{Status: 400, Code: "WRONG_CREDENTIALS", Message: "Wrong username or password"},
			},
			want: view.ErrorResponse{
				Status: 400,
				Code:   "WRONG_CREDENTIALS",
				Err:    "Wrong username or password",
			},
		},
		"valid stack error": {
			args: args{
				err: errors.WithStack(model.NewError(400, "bad_request", "bad request")),
			},
			want: view.ErrorResponse{
				Status: 400,
				Code:   "bad_request",
				Err:    "bad request",
			},
		},
		"valid viewmodel error": {
			args: args{
				err: view.ErrorResponse{
					Status: 400,
					Code:   "bad_request",
					Err:    "bad request",
				},
			},
			want: view.ErrorResponse{
				Status: 400,
				Code:   "bad_request",
				Err:    "bad request",
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tryParseError(tt.args.err)
			assert.Equal(t, tt.want, got)
		})
	}
}
