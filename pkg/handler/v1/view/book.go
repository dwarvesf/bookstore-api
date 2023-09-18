package view

// Book represent the book
type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Topic  *Topic `json:"topic"`
} // @name Book

// GetBooksRequest represent the get books response
type GetBooksRequest struct {
	Page     int    `form:"page" validate:"required"`
	PageSize int    `form:"pageSize" validate:"required"`
	Sort     string `form:"sort"`
	Query    string `form:"query"`
	TopicID  int    `form:"topicID"`
} // @name GetBooksRequest

// BooksResponse represent the get books response
type BooksResponse = ListResponse[Book] // @name BooksResponse
