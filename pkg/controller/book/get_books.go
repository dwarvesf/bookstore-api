package book

import (
	"context"

	"github.com/dwarvesf/bookstore-api/pkg/middleware"
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/repository/db"
)

// GetBooks get books
func (c *impl) GetBooks(ctx context.Context, q model.ListQuery, topicID int) (*model.ListResult[model.Book], error) {
	const spanName = "GetBooks"
	ctx, span := c.monitor.Start(ctx, spanName)
	defer span.End()

	uID, err := middleware.UserIDFromContext(ctx)
	if err != nil {
		return nil, model.ErrInvalidToken
	}

	dbCtx := db.FromContext(ctx)
	books, err := c.repo.Book.GetList(dbCtx, q, topicID, uID)
	if err != nil {
		return nil, err
	}

	return books, nil
}
