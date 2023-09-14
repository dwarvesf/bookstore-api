package view

// Topic represent the topic
type Topic struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
} // @name Topic

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
	PageSize int    `form:"page_size" validate:"required"`
	Sort     string `form:"sort"`
	Query    string `form:"query"`
	TopicID  int    `form:"topic_id"`
} // @name GetBooksRequest

// GetBooksResponse represent the get books response
type GetBooksResponse struct {
	Metadata Metadata `json:"metadata"`
	Data     []*Book  `json:"data"`
} // @name GetBooksResponse
