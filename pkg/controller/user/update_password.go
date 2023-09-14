package user

import (
	"context"

	"github.com/dwarvesf/df-bookstore-srv/pkg/middleware"
	"github.com/dwarvesf/df-bookstore-srv/pkg/model"
	"github.com/dwarvesf/df-bookstore-srv/pkg/repository/db"
)

func (c impl) UpdatePassword(ctx context.Context, user model.UpdatePasswordRequest) error {
	const spanName = "UpdatePasswordController"
	ctx, span := c.monitor.Start(ctx, spanName)
	defer span.End()

	uID, err := middleware.UserIDFromContext(ctx)
	if err != nil {
		return model.ErrInvalidToken
	}
	dbCtx := db.FromContext(ctx)
	u, err := c.repo.User.GetByID(dbCtx, uID)
	if err != nil {
		return err
	}

	if u.HashedPassword != user.OldPassword {
		return model.ErrInvalidCredentials
	}

	err = c.repo.User.UpdatePassword(dbCtx, uID, user.NewPassword)
	return err
}
