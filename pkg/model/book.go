package model

// UpdateBookRequest represent the update book request
type UpdateBookRequest struct {
	ID      int
	Name    *string
	Author  *string
	TopicID *int
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
	UserID  int
}
