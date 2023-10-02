package user

import (
	"context"

	"github.com/dwarvesf/bookstore-api/pkg/middleware"
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/repository/db"
)

func (c *impl) UpdateUser(ctx context.Context, user model.UpdateUserRequest) (*model.User, error) {
	const spanName = "UpdateUserController"
	ctx, span := c.monitor.Start(ctx, spanName)
	defer span.End()

	uID, err := middleware.UserIDFromContext(ctx)
	if err != nil {
		return nil, model.ErrInvalidToken
	}

	dbCtx := db.FromContext(ctx)
	_, err = c.repo.User.GetByID(dbCtx, uID)
	if err != nil {
		return nil, err
	}

	updated, err := c.repo.User.Update(dbCtx, uID, user)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
