package portal

import (
	"github.com/dwarvesf/bookstore-api/pkg/config"
	"github.com/dwarvesf/bookstore-api/pkg/controller/auth"
	"github.com/dwarvesf/bookstore-api/pkg/controller/user"
	"github.com/dwarvesf/bookstore-api/pkg/logger"
	"github.com/dwarvesf/bookstore-api/pkg/logger/monitor"
	"github.com/dwarvesf/bookstore-api/pkg/repository"
	"github.com/dwarvesf/bookstore-api/pkg/service"
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
