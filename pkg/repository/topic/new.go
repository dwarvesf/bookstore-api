package topic

import (
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/repository/db"
	"github.com/dwarvesf/bookstore-api/pkg/repository/orm"
)

// Repo represent the book
type Repo interface {
	IsExist(ctx db.Context, ID int) (bool, error)
}

// New return new book repo
func New() Repo {
	return &repo{}
}

func toTopicModel(topic *orm.Topic) *model.Topic {
	if topic == nil {
		return nil
	}
	return &model.Topic{
		ID:   topic.ID,
		Name: topic.Name,
		Code: topic.Code,
	}
}
