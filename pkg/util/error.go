package util

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dwarvesf/bookstore-api/pkg/handler/v1/view"
	"github.com/dwarvesf/bookstore-api/pkg/logger/monitor"
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// HandleError handle the rest error
func HandleError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if ok {
		errStr := ""
		for _, e := range errs {
			switch e.Tag() {
			case "password":
				errStr = fmt.Sprintf("%sPassword must contain at least 1 uppercase, 1 lowercase, 1 special character and length must be at least 8 characters; ", errStr)
			case "author":
				errStr = fmt.Sprintf("%sAuthor just only contain alphabet characters and space; ", errStr)
			case "required":
				errStr = fmt.Sprintf("%sField %s is required; ", errStr, e.Field())
			default:
				errStr = fmt.Sprintf("%s%s; ", errStr, e)
			}
		}

		err = view.ErrBadRequest(errors.New(errStr))
	}

	e := tryParseError(err)
	c.JSON(e.Status, gin.H{
		"status":  e.Status,
		"code":    e.Code,
		"message": e.Err,
		"traceID": monitor.GetTraceID(c.Request.Context()),
	})
}

func tryParseError(err error) view.ErrorResponse {
	var e model.Error
	ok := errors.As(err, &e)
	if ok {
		return view.ErrorResponse{
			Status: e.Status,
			Code:   e.Code,
			Err:    e.Message,
		}
	}

	var viewErr view.ErrorResponse
	ok = errors.As(err, &viewErr)
	if ok {
		return viewErr
	}

	return view.ErrorResponse{
		Status: http.StatusInternalServerError,
		Code:   "INTERNAL_ERROR",
		Err:    err.Error(),
	}
}
