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
// @Param pageSize query int false "Page Size"
// @Param sort query string false "Sort"
// @Param query query string false "Query"
// @Param topicID query int false "Topic ID"
// @Success 200 {object} BooksResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /books [get]
func (h Handler) GetBooks(c *gin.Context) {
	const spanName = "GetBooks"
	ctx, span := h.monitor.Start(c.Request.Context(), spanName)
	defer span.End()

	var req view.GetBooksRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		util.HandleError(c, err)
		return
	}

	rs, err := h.bookCtrl.GetBooks(ctx, model.ListQuery{
		Page:     req.Page,
		PageSize: req.PageSize,
		Sort:     req.Sort,
		Query:    req.Query,
	}, req.TopicID)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	books := make([]view.Book, 0, len(rs.Data))
	for _, b := range rs.Data {
		book := toBookView(&b)

		if book != nil {
			books = append(books, *book)
		}
	}

	c.JSON(http.StatusOK, view.BooksResponse{
		Metadata: view.Metadata{
			Page:         rs.Pagination.Page,
			PageSize:     rs.Pagination.PageSize,
			TotalPages:   rs.Pagination.TotalPages,
			TotalRecords: rs.Pagination.TotalRecords,
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
// @Router /books [post]
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
