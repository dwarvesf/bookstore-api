package book

import (
	"context"

	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/repository/db"
)

// CreateBook get books
func (c *impl) CreateBook(ctx context.Context, book model.CreateBookRequest) (*model.Book, error) {
	const spanName = "CreateBook"
	ctx, span := c.monitor.Start(ctx, spanName)
	defer span.End()

	dbCtx := db.FromContext(ctx)
	// validate topic
	isExist, err := c.repo.Topic.IsExist(dbCtx, book.TopicID)
	if err != nil {
		return nil, err
	}

	if !isExist {
		return nil, model.ErrTopicNotFound
	}

	res, err := c.repo.Book.Create(dbCtx, book)
	if err != nil {
		return nil, err
	}

	return res, nil
}
