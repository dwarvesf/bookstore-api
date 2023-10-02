package view

// Book represent the book
type Book struct {
	ID     int    `json:"id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Author string `json:"author"`
	Topic  *Topic `json:"topic"`
} // @name Book

// GetBooksRequest represent the get books response
type GetBooksRequest struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Sort     string `form:"sort"`
	Query    string `form:"query"`
	TopicID  int    `form:"topicId"`
} // @name GetBooksRequest

// BooksResponse represent the get books response
type BooksResponse = ListResponse[Book] // @name BooksResponse

// CreateBookRequest represent the create book request
type CreateBookRequest struct {
	Name    string `json:"name" validate:"required" binding:"required"`
	Author  string `json:"author" validate:"required" binding:"required,author"`
	TopicID int    `json:"topicId" validate:"required" binding:"required,gt=0"`
} // @name CreateBookRequest

// UpdateBookRequest represent the update book request
type UpdateBookRequest struct {
	Name    *string `json:"name"`
	Author  *string `json:"author" binding:"omitempty,author"`
	TopicID *int    `json:"topicId" binding:"omitempty,gt=0"`
} // @name UpdateBookRequest

// BookResponse represent the book response
type BookResponse = Response[Book] // @name BookResponse
