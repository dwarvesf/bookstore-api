package model

// UpdateBookRequest represent the update book request
type UpdateBookRequest struct {
	Name    string
	Author  string
	TopicID int
}

// Topic represent the topic
type Topic struct {
	ID   int
	Name string
	Code string
}

// Book represent the book
type Book struct {
	ID     int
	Name   string
	Author string
	Topic  *Topic
}

// CreateBookRequest represent the create book request
type CreateBookRequest struct {
	Name    string
	Author  string
	TopicID int
}
