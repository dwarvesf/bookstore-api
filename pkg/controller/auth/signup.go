package auth

import (
	"context"

	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/repository/db"
	"github.com/pkg/errors"
)

func (c impl) Signup(ctx context.Context, req model.SignupRequest) (err error) {
	const spanName = "SignupController"
	ctx, span := c.monitor.Start(ctx, spanName)
	defer span.End()

	req.Salt = c.passwordHelper.GenerateSalt()
	hashedPassword, err := c.passwordHelper.Hash(req.Password, req.Salt)
	if err != nil {
		return err
	}
	req.HashedPassword = hashedPassword
	req.Role = model.RoleUser
	req.Status = model.StatusActive

	dbCtx, finalFn := db.NewTransaction(ctx)
	defer finalFn(err)

	//  check if email is existed
	_, err = c.repo.User.GetByEmail(dbCtx, req.Email)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return err
	}

	if err == nil {
		return model.ErrEmailExisted
	}

	_, err = c.repo.User.Create(dbCtx, req)
	if err != nil {
		return err
	}

	return nil
}
