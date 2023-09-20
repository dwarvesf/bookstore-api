package book

import (
	"context"

	"github.com/dwarvesf/bookstore-api/pkg/middleware"
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/repository/db"
)

// DeleteBook delete book
func (c *impl) DeleteBook(ctx context.Context, ID int) error {
	const spanName = "DeleteBook"
	ctx, span := c.monitor.Start(ctx, spanName)
	defer span.End()

	dbCtx := db.FromContext(ctx)
	userID, err := middleware.UserIDFromContext(ctx)
	if err != nil {
		return model.ErrInvalidToken
	}

	isExist, err := c.repo.Book.IsExist(dbCtx, ID, userID)
	if err != nil {
		return err
	}

	if !isExist {
		return model.ErrBookNotFound
	}

	return c.repo.Book.Delete(dbCtx, ID)
}
