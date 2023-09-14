package portal

import (
	"github.com/dwarvesf/df-bookstore-srv/pkg/config"
	"github.com/dwarvesf/df-bookstore-srv/pkg/controller/auth"
	"github.com/dwarvesf/df-bookstore-srv/pkg/controller/user"
	"github.com/dwarvesf/df-bookstore-srv/pkg/logger"
	"github.com/dwarvesf/df-bookstore-srv/pkg/logger/monitor"
	"github.com/dwarvesf/df-bookstore-srv/pkg/repository"
	"github.com/dwarvesf/df-bookstore-srv/pkg/service"
)

// Handler for app
type Handler struct {
	cfg      config.Config
	log      logger.Log
	svc      service.Service
	monitor  monitor.Tracer
	authCtrl auth.Controller
	userCtrl user.Controller
}

// New will return an instance of Auth struct
func New(cfg config.Config, l logger.Log, repo *repository.Repo, svc service.Service, monitor monitor.Tracer) *Handler {
	return &Handler{
		cfg:      cfg,
		log:      l,
		svc:      svc,
		monitor:  monitor,
		authCtrl: auth.NewAuthController(cfg, repo, monitor),
		userCtrl: user.NewUserController(cfg, repo, monitor),
	}
}
