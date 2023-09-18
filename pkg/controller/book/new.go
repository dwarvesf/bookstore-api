package book

import (
	"context"

	"github.com/dwarvesf/bookstore-api/pkg/config"
	"github.com/dwarvesf/bookstore-api/pkg/logger/monitor"
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/repository"
)

// Controller auth controller
type Controller interface {
	GetBooks(ctx context.Context, q model.ListQuery, topicID int) (*model.ListResult[model.Book], error)
}

type impl struct {
	repo    *repository.Repo
	cfg     config.Config
	monitor monitor.Tracer
}

// NewBookController new book controller
func NewBookController(cfg config.Config, r *repository.Repo, monitor monitor.Tracer) Controller {
	return &impl{
		repo:    r,
		cfg:     cfg,
		monitor: monitor,
	}
}
