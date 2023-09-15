package book

import (
	"context"

	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/repository/db"
)

// UpdateBook update book
func (c *impl) UpdateBook(ctx context.Context, book model.UpdateBookRequest) (*model.Book, error) {
	const spanName = "UpdateBook"
	ctx, span := c.monitor.Start(ctx, spanName)
	defer span.End()

	dbCtx := db.FromContext(ctx)
	// validate topic
	if book.TopicID != 0 {
		isExist, err := c.repo.Topic.IsExist(dbCtx, book.TopicID)
		if err != nil {
			return nil, err
		}

		if !isExist {
			return nil, model.ErrTopicNotFound
		}
	}

	_, err := c.repo.Book.Update(dbCtx, book)
	if err != nil {
		return nil, err
	}

	return c.repo.Book.GetByID(dbCtx, book.ID)
}
