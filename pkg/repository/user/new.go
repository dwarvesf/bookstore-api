package user

import (
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/repository/db"
	"github.com/dwarvesf/bookstore-api/pkg/repository/orm"
)

// Repo represent the user
type Repo interface {
	GetByID(ctx db.Context, uID int) (*model.User, error)
	GetList(ctx db.Context, q model.ListQuery) (*model.ListResult[model.User], error)
	Count(ctx db.Context) (int64, error)
	GetByEmail(ctx db.Context, email string) (*model.User, error)
	Create(ctx db.Context, user model.SignupRequest) (*model.User, error)
	Update(ctx db.Context, uID int, user model.UpdateUserRequest) (*model.User, error)
	UpdatePassword(ctx db.Context, uID int, newPassword string) error
	IsExist(ctx db.Context, ID int) (bool, error)
}

// New return new user repo
func New() Repo {
	return &repo{}
}

func toUserModel(user *orm.User) *model.User {
	if user == nil {
		return nil
	}
	return &model.User{
		ID:             user.ID,
		Email:          user.Email,
		FullName:       user.Name,
		Status:         user.Status,
		Avatar:         user.Avatar,
		HashedPassword: user.HashedPassword,
		Role:           user.Role,
		Salt:           user.Salt,
	}
}
