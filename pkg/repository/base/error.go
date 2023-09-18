package base

import (
	"database/sql"
	"errors"

	"github.com/dwarvesf/bookstore-api/pkg/model"
)

// GetOneErrorHandler handle error when get one
func GetOneErrorHandler(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return model.ErrNotFound
	}
	return err

}
