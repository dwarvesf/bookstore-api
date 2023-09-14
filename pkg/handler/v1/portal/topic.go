package portal

import (
	"github.com/dwarvesf/bookstore-api/pkg/handler/v1/view"
	"github.com/dwarvesf/bookstore-api/pkg/model"
)

func toTopicView(topic *model.Topic) *view.Topic {
	if topic == nil {
		return nil
	}

	return &view.Topic{
		ID:   topic.ID,
		Name: topic.Name,
		Code: topic.Code,
	}
}
