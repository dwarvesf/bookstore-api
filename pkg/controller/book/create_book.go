package book

import (
	"context"

	"github.com/dwarvesf/bookstore-api/pkg/middleware"
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/repository/db"
)

// CreateBook create new book
func (c *impl) CreateBook(ctx context.Context, book model.CreateBookRequest) (*model.Book, error) {
	const spanName = "CreateBook"
	ctx, span := c.monitor.Start(ctx, spanName)
	defer span.End()

	uID, err := middleware.UserIDFromContext(ctx)
	if err != nil {
		return nil, model.ErrInvalidToken
	}
	book.UserID = uID

	dbCtx := db.FromContext(ctx)

	if book.TopicID > 0 {
		// validate topic
		isExist, err := c.repo.Topic.IsExist(dbCtx, book.TopicID)
		if err != nil {
			return nil, err
		}

		if !isExist {
			return nil, model.ErrTopicNotFound
		}
	}

	newBook, err := c.repo.Book.Create(dbCtx, book)
	if err != nil {
		return nil, err
	}

	return c.repo.Book.GetByID(dbCtx, newBook.ID)
}
