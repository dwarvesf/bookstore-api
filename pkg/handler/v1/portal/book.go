package portal

import (
	"net/http"

	"github.com/dwarvesf/bookstore-api/pkg/handler/v1/view"
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/util"
	"github.com/gin-gonic/gin"
)

// GetBooks godoc
// @Summary Get list of books
// @Description Get list of books
// @Tags Book
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param page query int false "Page"
// @Param page_size query int false "Page Size"
// @Param sort query string false "Sort"
// @Param query query string false "Query"
// @Param topic_id query int false "Topic ID"
// @Success 200 {object} GetBooksResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /portal/books [get]
func (h Handler) GetBooks(c *gin.Context) {
	const spanName = "GetBooks"
	newCtx, span := h.monitor.Start(c.Request.Context(), spanName)
	defer span.End()

	var req view.GetBooksRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		util.HandleError(c, err)
		return
	}

	var types map[string]interface{}
	if req.TopicID != 0 {
		types = map[string]interface{}{
			"topic_id": req.TopicID,
		}
	}

	rs, err := h.bookCtrl.GetBooks(newCtx, model.ListQuery{
		Page:     req.Page,
		PageSize: req.PageSize,
		Sort:     req.Sort,
		Query:    req.Query,
		Types:    types,
	})
	if err != nil {
		util.HandleError(c, err)
		return
	}

	books := make([]*view.Book, 0, len(rs.Data))
	for _, b := range rs.Data {
		books = append(books, toBookView(&b))
	}

	c.JSON(http.StatusOK, view.GetBooksResponse{
		Metadata: view.Metadata{
			Page:         rs.Pagination.Page,
			PageSize:     rs.Pagination.PageSize,
			TotalPages:   rs.Pagination.TotalPages,
			TotalRecords: rs.Pagination.TotalRecords,
			Sort:         rs.Pagination.Sort,
		},
		Data: books,
	})
}

func toBookView(b *model.Book) *view.Book {
	if b == nil {
		return nil
	}

	return &view.Book{
		ID:     b.ID,
		Name:   b.Name,
		Author: b.Author,
		Topic:  toTopicView(b.Topic),
	}
}

// CreateBook godoc
// @Summary Create new book
// @Description Create new book
// @Tags Book
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param body body CreateBookRequest true "Create Book Request"
// @Success 200 {object} Book
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /portal/books [post]
func (h Handler) CreateBook(c *gin.Context) {
	const spanName = "CreateBook"
	newCtx, span := h.monitor.Start(c.Request.Context(), spanName)
	defer span.End()

	var req view.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.HandleError(c, err)
		return
	}

	rs, err := h.bookCtrl.CreateBook(newCtx, model.CreateBookRequest{
		Name:    req.Name,
		Author:  req.Author,
		TopicID: req.TopicID,
	})
	if err != nil {
		util.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, toBookView(rs))
}
