package portal

import (
	"net/http"
	"strconv"

	"github.com/dwarvesf/bookstore-api/pkg/handler/v1/view"
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/util"

	"github.com/gin-gonic/gin"
)

// GetBooks godoc
// @Summary Get list of books
// @Description Get list of books
// @id getBooks
// @Tags Book
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param page query int false "Page"
// @Param pageSize query int false "Page Size"
// @Param sort query string false "Sort"
// @Param query query string false "Query"
// @Param topicId query int false "Topic ID"
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

// GetBook godoc
// @Summary Get book by id
// @Description Get book by id
// @id getBook
// @Tags Book
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path int true "Book ID"
// @Success 200 {object} BookResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /books/{id} [get]
func (h Handler) GetBook(c *gin.Context) {
	const spanName = "GetBook"
	ctx, span := h.monitor.Start(c.Request.Context(), spanName)
	defer span.End()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.HandleError(c, view.ErrBadRequest(model.ErrInvalidBookID))
		return
	}

	rs, err := h.bookCtrl.GetBook(ctx, id)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, view.BookResponse{Data: *toBookView(rs)})
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
// @id createBook
// @Tags Book
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param body body CreateBookRequest true "Create Book Request"
// @Success 200 {object} BookResponse
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

	c.JSON(http.StatusOK, view.BookResponse{Data: *toBookView(rs)})
}

// UpdateBook godoc
// @Summary Update book
// @Description Update book
// @id updateBook
// @Tags Book
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path int true "Book ID"
// @Param body body UpdateBookRequest true "Update Book Request"
// @Success 200 {object} BookResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /books/{id} [put]
func (h Handler) UpdateBook(c *gin.Context) {
	const spanName = "UpdateBook"
	newCtx, span := h.monitor.Start(c.Request.Context(), spanName)
	defer span.End()

	bookID := c.Param("id")
	ID, err := strconv.Atoi(bookID)
	if err != nil {
		util.HandleError(c, model.ErrInvalidBookID)
		return
	}

	var req view.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.HandleError(c, err)
		return
	}

	res, err := h.bookCtrl.UpdateBook(newCtx, model.UpdateBookRequest{
		ID:      ID,
		Name:    req.Name,
		Author:  req.Author,
		TopicID: req.TopicID,
	})
	if err != nil {
		util.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, view.BookResponse{Data: *toBookView(res)})
}

// DeleteBook godoc
// @Summary Delete book by id
// @Description Delete book by id
// @id deleteBook
// @Tags Book
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path int true "Book ID"
// @Success 200 {object} MessageResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /books/{id} [delete]
func (h Handler) DeleteBook(c *gin.Context) {
	const spanName = "DeleteBook"
	newCtx, span := h.monitor.Start(c.Request.Context(), spanName)
	defer span.End()

	bookID := c.Param("id")
	ID, err := strconv.Atoi(bookID)
	if err != nil {
		util.HandleError(c, model.ErrInvalidBookID)
		return
	}

	err = h.bookCtrl.DeleteBook(newCtx, ID)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, view.MessageResponse{Data: view.Message{Message: "OK"}})
}
