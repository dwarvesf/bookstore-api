package model

import "net/http"

var (
	// ErrNoAuthHeader is the error for no authorization header
	ErrNoAuthHeader = Error{
		Status:  http.StatusUnauthorized,
		Code:    "Unauthorized",
		Message: "No authorization header",
	}
	// ErrInvalidToken is the error for invalid token
	ErrInvalidToken = Error{
		Status:  http.StatusUnauthorized,
		Code:    "Unauthorized",
		Message: "Unauthorized",
	}

	// ErrUnexpectedAuthorizationHeader is the error for unexpected authorization header
	ErrUnexpectedAuthorizationHeader = Error{
		Status:  http.StatusUnauthorized,
		Code:    "Unauthorized",
		Message: "Unexpected authorization headers",
	}

	// ErrInvalidCredentials is the error for invalid credentials
	ErrInvalidCredentials = Error{
		Status:  http.StatusBadRequest,
		Code:    "WRONG_CREDENTIALS",
		Message: "Wrong username or password",
	}

	// ErrNotFound is the error for not found
	ErrNotFound = Error{
		Status:  http.StatusNotFound,
		Code:    "NOT_FOUND",
		Message: "not found",
	}

	// ErrEmailExisted is the error for email existed
	ErrEmailExisted = Error{
		Status:  http.StatusBadRequest,
		Code:    "EMAIL_EXISTED",
		Message: "email existed",
	}

	// ErrTopicNotFound is the error for topic not found
	ErrTopicNotFound = Error{
		Status:  http.StatusNotFound,
		Code:    "TOPIC_NOT_FOUND",
		Message: "topic not found",
	}

	// ErrUserNotFound is the error for user not found
	ErrUserNotFound = Error{
		Status:  http.StatusNotFound,
		Code:    "USER_NOT_FOUND",
		Message: "user not found",
	}

	// ErrInvalidBookID is the error for invalid book id
	ErrInvalidBookID = Error{
		Status:  http.StatusBadRequest,
		Code:    "INVALID_BOOK_ID",
		Message: "invalid book id",
	}

	// ErrBookNotFound is the error for book not found
	ErrBookNotFound = Error{
		Status:  http.StatusNotFound,
		Code:    "BOOK_NOT_FOUND",
		Message: "book not found",
	}
)

// Error in server
type Error struct {
	Status  int
	Code    string
	Message string
}

func (e Error) Error() string {
	return e.Message
}

// NewError new a error with message
func NewError(status int, code, msg string) error {
	return Error{
		Status:  status,
		Code:    code,
		Message: msg,
	}
}
