package user

import (
	"context"

	"github.com/dwarvesf/bookstore-api/pkg/middleware"
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/repository/db"
)

func (c *impl) Me(ctx context.Context) (*model.User, error) {
	const spanName = "MeController"
	ctx, span := c.monitor.Start(ctx, spanName)
	defer span.End()

	uID, err := middleware.UserIDFromContext(ctx)
	if err != nil {
		return nil, model.ErrInvalidToken
	}

	dbCtx := db.FromContext(ctx)

	u, err := c.repo.User.GetByID(dbCtx, uID)
	if err != nil {
		return nil, err
	}

	return u, nil
}
