package portal

import (
	"net/http"

	"github.com/dwarvesf/bookstore-api/pkg/handler/v1/view"
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/util"
	"github.com/gin-gonic/gin"
)

// GetTopics godoc
// @Summary Get all topics
// @Description Get all topics
// @id getTopics
// @Tags Topic
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} TopicsResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /topics [get]
func (h Handler) GetTopics(c *gin.Context) {
	const spanName = "GetTopics"
	newCtx, span := h.monitor.Start(c.Request.Context(), spanName)
	defer span.End()

	rs, err := h.topicCtrl.GetTopics(newCtx)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	books := make([]view.Topic, 0, len(rs))
	for _, b := range rs {
		newBook := toTopicView(&b)
		if newBook != nil {
			books = append(books, *newBook)
		}
	}

	c.JSON(http.StatusOK, view.TopicsResponse{
		Data: books,
	})
}

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
