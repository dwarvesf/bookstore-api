package auth

import (
	"context"

	"github.com/dwarvesf/bookstore-api/pkg/config"
	"github.com/dwarvesf/bookstore-api/pkg/logger/monitor"
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/repository"
	"github.com/dwarvesf/bookstore-api/pkg/service/jwthelper"
	"github.com/dwarvesf/bookstore-api/pkg/service/passwordhelper"
)

// Controller auth controller
type Controller interface {
	Login(ctx context.Context, req model.LoginRequest) (*model.LoginResponse, error)
	Signup(ctx context.Context, req model.SignupRequest) error
}

type impl struct {
	repo           *repository.Repo
	jwtHelper      jwthelper.Helper
	cfg            config.Config
	monitor        monitor.Tracer
	passwordHelper passwordhelper.Helper
}

// NewAuthController new auth controller
func NewAuthController(cfg config.Config, r *repository.Repo, monitor monitor.Tracer) Controller {
	return &impl{
		repo:           r,
		jwtHelper:      jwthelper.NewHelper(cfg.SecretKey),
		cfg:            cfg,
		monitor:        monitor,
		passwordHelper: passwordhelper.NewScrypt(),
	}
}
