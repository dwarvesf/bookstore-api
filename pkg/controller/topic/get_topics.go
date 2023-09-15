package topic

import (
	"context"

	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/repository/db"
)

// GetTopics get topics
func (c *impl) GetTopics(ctx context.Context) ([]*model.Topic, error) {
	const spanName = "GetTopics"
	ctx, span := c.monitor.Start(ctx, spanName)
	defer span.End()

	dbCtx := db.FromContext(ctx)
	books, err := c.repo.Topic.GetAll(dbCtx)
	if err != nil {
		return nil, err
	}

	return books, nil
}
