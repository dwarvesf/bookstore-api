package topic

import (
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/repository/db"
	"github.com/dwarvesf/bookstore-api/pkg/repository/orm"
)

type repo struct {
}

// IsExist check if topic is exist
func (r *repo) IsExist(ctx db.Context, ID int) (bool, error) {
	return orm.TopicExists(ctx, ctx.DB, ID)
}

// GetAll get all topics
func (r *repo) GetAll(ctx db.Context) ([]model.Topic, error) {
	topics, err := orm.Topics().All(ctx.Context, ctx.DB)
	if err != nil {
		return nil, err
	}

	rs := []model.Topic{}
	for _, topic := range topics {
		topicModel := toTopicModel(topic)

		if topicModel != nil {
			rs = append(rs, *topicModel)
		}
	}

	return rs, nil
}
