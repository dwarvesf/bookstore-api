package book

import (
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/repository/db"
	"github.com/dwarvesf/bookstore-api/pkg/repository/orm"
)

// Repo represent the book
type Repo interface {
	GetByID(ctx db.Context, uID int) (*model.Book, error)
	GetList(ctx db.Context, q model.ListQuery) (*model.ListResult[model.Book], error)
	Count(ctx db.Context) (int64, error)
	Create(ctx db.Context, book model.CreateBookRequest) (*model.Book, error)
	Update(ctx db.Context, uID int, book model.UpdateBookRequest) (*model.Book, error)
}

// New return new book repo
func New() Repo {
	return &repo{}
}

func toBookModel(book *orm.Book) *model.Book {
	if book == nil {
		return nil
	}

	var topic *model.Topic
	if book.R != nil {
		topic = toTopicModel(book.R.Topic)
	}

	return &model.Book{
		ID:     book.ID,
		Name:   book.Name,
		Author: book.Author.String,
		Topic:  topic,
	}
}

func toTopicModel(topic *orm.Topic) *model.Topic {
	if topic == nil {
		return nil
	}
	return &model.Topic{
		ID:   topic.ID,
		Name: topic.Name,
		Code: topic.Code,
	}
}
