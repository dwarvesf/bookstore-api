package repository

import (
	"github.com/dwarvesf/bookstore-api/pkg/repository/book"
	"github.com/dwarvesf/bookstore-api/pkg/repository/topic"
	"github.com/dwarvesf/bookstore-api/pkg/repository/user"
)

// Repo represent the repository
type Repo struct {
	User  user.Repo
	Book  book.Repo
	Topic topic.Repo
}

// NewRepo will create an object that represent the Repo interface
func NewRepo() *Repo {
	return &Repo{
		User:  user.New(),
		Book:  book.New(),
		Topic: topic.New(),
	}
}
