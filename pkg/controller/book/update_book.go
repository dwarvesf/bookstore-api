package book

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dwarvesf/bookstore-api/pkg/middleware"
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/repository/db"
)

// UpdateBook update book
func (c *impl) UpdateBook(ctx context.Context, book model.UpdateBookRequest) (*model.Book, error) {
	const spanName = "UpdateBook"
	ctx, span := c.monitor.Start(ctx, spanName)
	defer span.End()

	dbCtx := db.FromContext(ctx)

	uID, err := middleware.UserIDFromContext(ctx)
	if err != nil {
		return nil, model.ErrInvalidToken
	}

	_, err = c.repo.Book.GetByUserAndID(dbCtx, uID, book.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrBookNotFound
		}

		return nil, err
	}

	if book.TopicID > 0 {
		isExist, err := c.repo.Topic.IsExist(dbCtx, book.TopicID)
		if err != nil {
			return nil, err
		}

		if !isExist {
			return nil, model.ErrTopicNotFound
		}
	}

	_, err = c.repo.Book.Update(dbCtx, book)
	if err != nil {
		return nil, err
	}

	return c.repo.Book.GetByID(dbCtx, book.ID)
}
