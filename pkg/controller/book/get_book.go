package book

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dwarvesf/bookstore-api/pkg/middleware"
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/repository/db"
)

// GetBook get book by id
func (c *impl) GetBook(ctx context.Context, id int) (*model.Book, error) {
	const spanName = "GetBook"
	ctx, span := c.monitor.Start(ctx, spanName)
	defer span.End()

	uID, err := middleware.UserIDFromContext(ctx)
	if err != nil {
		return nil, model.ErrInvalidToken
	}

	dbCtx := db.FromContext(ctx)
	rs, err := c.repo.Book.GetByUserAndID(dbCtx, uID, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, model.ErrBookNotFound
	}

	if err != nil {
		return nil, err
	}

	return rs, nil
}
