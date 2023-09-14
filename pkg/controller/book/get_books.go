package book

import (
	"context"

	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/repository/db"
)

// GetBooks get books
func (c *impl) GetBooks(ctx context.Context, q model.ListQuery) (*model.ListResult[model.Book], error) {
	const spanName = "GetBooks"
	ctx, span := c.monitor.Start(ctx, spanName)
	defer span.End()

	dbCtx := db.FromContext(ctx)
	books, err := c.repo.Book.GetList(dbCtx, q)
	if err != nil {
		return nil, err
	}

	return books, nil
}
