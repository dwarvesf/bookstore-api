package book

import (
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/repository/base"
	"github.com/dwarvesf/bookstore-api/pkg/repository/db"
	"github.com/dwarvesf/bookstore-api/pkg/repository/orm"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type repo struct {
}

func (r *repo) GetList(ctx db.Context, q model.ListQuery, topicID int, uID int) (*model.ListResult[model.Book], error) {
	fnSet := base.GetListFuncSet[orm.Book, model.Book]{
		PrepareQueryFn: func(ctx db.Context, q model.ListQuery) []qm.QueryMod {
			queryParams := []qm.QueryMod{}
			if q.Query != "" {
				queryParams = append(queryParams, qm.Where("user_id=? AND (lower(name) LIKE lower(?) OR lower(author) LIKE lower(?))", uID, "%"+q.Query+"%", "%"+q.Query+"%"))
			}

			if topicID != 0 {
				queryParams = append(queryParams, qm.Where("topic_id=?", topicID))
			}
			return queryParams
		},
		CounableFn: func(q []qm.QueryMod) base.Counable {
			return orm.Books(q...)
		},
		QueryListFn: func(q []qm.QueryMod) ([]*orm.Book, error) {
			q = append(q, qm.Load("Topic"))
			return orm.Books(q...).All(ctx.Context, ctx.DB)
		},
		MappingFn: toBookModel,
	}

	return base.GetList(ctx, q, fnSet)
}

func (r *repo) Count(ctx db.Context) (int64, error) {
	return orm.Books().Count(ctx.Context, ctx.DB)
}

func (r *repo) GetByID(ctx db.Context, uID int) (*model.Book, error) {
	dt, err := orm.Books(qm.Where("id=?", uID), qm.Load("Topic")).One(ctx.Context, ctx.DB)
	return toBookModel(dt), err
}

func (r *repo) Create(ctx db.Context, book model.CreateBookRequest) (*model.Book, error) {
	u := &orm.Book{
		Name:    book.Name,
		Author:  null.String{String: book.Author, Valid: book.Author != ""},
		TopicID: null.Int{Int: book.TopicID, Valid: book.TopicID != 0},
		UserID:  book.UserID,
	}

	err := u.Insert(ctx, ctx.DB, boil.Infer())

	u, err = orm.Books(qm.Where("id=?", u.ID), qm.Load("Topic")).One(ctx.Context, ctx.DB)
	return toBookModel(u), err
}

func (r *repo) Update(ctx db.Context, book model.UpdateBookRequest) (*model.Book, error) {
	u, err := orm.FindBook(ctx, ctx.DB, book.ID)
	if err != nil {
		return nil, err
	}

	u.Name = book.Name
	u.Author = null.String{String: book.Author, Valid: book.Author != ""}
	u.TopicID = null.Int{Int: book.TopicID, Valid: book.TopicID != 0}

	_, err = u.Update(ctx, ctx.DB, boil.Infer())
	return toBookModel(u), err
}

func (r *repo) IsExist(ctx db.Context, id int, userID int) (bool, error) {
	return orm.Books(qm.Where("id=? AND user_id=?", id, userID)).Exists(ctx.Context, ctx.DB)
}

func (r *repo) Delete(ctx db.Context, id int) error {
	_, err := orm.Books(qm.Where("id=?", id)).DeleteAll(ctx.Context, ctx.DB)
	return err
}
