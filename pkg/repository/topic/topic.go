package topic

import (
	"github.com/dwarvesf/bookstore-api/pkg/repository/db"
	"github.com/dwarvesf/bookstore-api/pkg/repository/orm"
)

type repo struct {
}

// IsExist check if topic is exist
func (r *repo) IsExist(ctx db.Context, ID int) (bool, error) {
	return orm.TopicExists(ctx, ctx.DB, ID)
}
