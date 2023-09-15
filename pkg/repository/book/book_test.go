package book

import (
	"reflect"
	"testing"

	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/repository/db"
	"github.com/dwarvesf/bookstore-api/pkg/repository/orm"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Test_repo_GetByID(t *testing.T) {
	db.WithTestingDB(t, func(ctx db.Context) {
		topic := &orm.Topic{
			Name: "topic1",
			Code: "code1",
		}
		err := topic.Insert(ctx, ctx.DB, boil.Infer())
		require.NoError(t, err)

		u := &orm.Book{
			Name:    "book1",
			Author:  null.String{String: "author1", Valid: true},
			TopicID: null.Int{Int: topic.ID, Valid: true},
		}
		err = u.Insert(ctx, ctx.DB, boil.Infer())
		require.NoError(t, err)

		type args struct {
			uID int
		}
		tests := map[string]struct {
			args    args
			want    *model.Book
			wantErr bool
		}{
			"success": {
				args: args{
					uID: u.ID,
				},
				want: &model.Book{
					ID:     u.ID,
					Name:   "book1",
					Author: "author1",
					Topic: &model.Topic{
						ID:   topic.ID,
						Name: "topic1",
						Code: "code1",
					},
				},
				wantErr: false,
			},
			"not found": {
				args: args{
					uID: u.ID + 1,
				},
				want:    nil,
				wantErr: true,
			},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				r := &repo{}
				got, err := r.GetByID(ctx, tt.args.uID)
				if (err != nil) != tt.wantErr {
					t.Errorf("repo.GetByID() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				require.Equal(t, tt.want, got)
			})
		}
	})
}

func Test_repo_Count(t *testing.T) {
	db.WithTestingDB(t, func(ctx db.Context) {
		topic := &orm.Topic{
			Name: "topic1",
			Code: "code1",
		}
		err := topic.Insert(ctx, ctx.DB, boil.Infer())
		require.NoError(t, err)

		u := &orm.Book{
			Name:    "book1",
			Author:  null.String{String: "author1", Valid: true},
			TopicID: null.Int{Int: topic.ID, Valid: true},
		}
		err = u.Insert(ctx, ctx.DB, boil.Infer())
		require.NoError(t, err)

		tests := map[string]struct {
			want    int64
			wantErr bool
		}{
			"success": {
				want:    1,
				wantErr: false,
			},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				r := &repo{}
				got, err := r.Count(ctx)
				if (err != nil) != tt.wantErr {
					t.Errorf("repo.Count() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("repo.Count() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func Test_repo_Create(t *testing.T) {
	db.WithTestingDB(t, func(ctx db.Context) {
		topic := &orm.Topic{
			Name: "topic1",
			Code: "code1",
		}
		err := topic.Insert(ctx, ctx.DB, boil.Infer())
		require.NoError(t, err)

		u := &orm.Book{
			Name:    "book1",
			Author:  null.String{String: "author1", Valid: true},
			TopicID: null.Int{Int: topic.ID, Valid: true},
		}
		err = u.Insert(ctx, ctx.DB, boil.Infer())
		require.NoError(t, err)

		type args struct {
			req model.CreateBookRequest
		}
		tests := map[string]struct {
			args    args
			want    *model.Book
			wantErr bool
		}{
			"success": {
				args: args{
					req: model.CreateBookRequest{
						Name:    "book2",
						Author:  "author2",
						TopicID: topic.ID,
					},
				},
				want: &model.Book{
					ID:     u.ID + 1,
					Name:   "book2",
					Author: "author2",
					Topic: &model.Topic{
						ID:   topic.ID,
						Name: "topic1",
						Code: "code1",
					},
				},
				wantErr: false,
			},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				r := &repo{}
				got, err := r.Create(ctx, tt.args.req)
				if (err != nil) != tt.wantErr {
					t.Errorf("repo.Create() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if !tt.wantErr {
					if !reflect.DeepEqual(got, tt.want) {
						t.Errorf("repo.Create() = %v, want %v", got, tt.want)
					}
				}
			})
		}
	})
}

func Test_repo_Update(t *testing.T) {
	db.WithTestingDB(t, func(ctx db.Context) {
		topic := &orm.Topic{
			Name: "topic1",
			Code: "code1",
		}
		err := topic.Insert(ctx, ctx.DB, boil.Infer())
		require.NoError(t, err)

		u := &orm.Book{
			Name:    "book1",
			Author:  null.String{String: "author1", Valid: true},
			TopicID: null.Int{Int: topic.ID, Valid: true},
		}
		err = u.Insert(ctx, ctx.DB, boil.Infer())
		require.NoError(t, err)

		type args struct {
			uID  int
			book model.UpdateBookRequest
		}
		tests := map[string]struct {
			args    args
			want    *model.Book
			wantErr bool
		}{
			"success": {
				args: args{
					uID: u.ID,
					book: model.UpdateBookRequest{
						Name:    "book2",
						Author:  "author2",
						TopicID: topic.ID,
					},
				},
				want: &model.Book{
					ID:     u.ID,
					Name:   "book2",
					Author: "author2",
				},
				wantErr: false,
			},
			"not found": {
				args: args{
					uID: u.ID + 1,
				},
				want:    nil,
				wantErr: true,
			},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				r := &repo{}
				got, err := r.Update(ctx, tt.args.uID, tt.args.book)
				if (err != nil) != tt.wantErr {
					t.Errorf("repo.Update() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("repo.Update() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func Test_repo_GetList(t *testing.T) {
	db.WithTestingDB(t, func(ctx db.Context) {
		topic := &orm.Topic{
			Name: "topic1",
			Code: "code1",
		}
		err := topic.Insert(ctx, ctx.DB, boil.Infer())
		require.NoError(t, err)

		u := &orm.Book{
			Name:    "book1",
			Author:  null.String{String: "author1", Valid: true},
			TopicID: null.Int{Int: topic.ID, Valid: true},
		}
		err = u.Insert(ctx, ctx.DB, boil.Infer())
		require.NoError(t, err)
		u = &orm.Book{
			Name:    "book2",
			Author:  null.String{String: "author2", Valid: true},
			TopicID: null.Int{Int: topic.ID, Valid: true},
		}
		err = u.Insert(ctx, ctx.DB, boil.Infer())
		require.NoError(t, err)

		type args struct {
			page     int
			pageSize int
			sort     string
			query    string
		}
		tests := map[string]struct {
			args    args
			want    *model.ListResult[model.Book]
			wantErr bool
		}{
			"success": {
				args: args{
					page:     1,
					pageSize: 10,
					sort:     "+id",
					query:    "",
				},
				want: &model.ListResult[model.Book]{
					Pagination: model.Pagination{
						Page:         1,
						PageSize:     10,
						TotalRecords: 2,
						TotalPages:   1,
						Offset:       0,
						Sort:         "id asc",
						HasNext:      false,
					},
					Data: []model.Book{
						{
							ID:     u.ID - 1,
							Name:   "book1",
							Author: "author1",
							Topic: &model.Topic{
								ID:   topic.ID,
								Name: "topic1",
								Code: "code1",
							},
						},
						{
							ID:     u.ID,
							Name:   "book2",
							Author: "author2",
							Topic: &model.Topic{
								ID:   topic.ID,
								Name: "topic1",
								Code: "code1",
							},
						},
					},
				},
				wantErr: false,
			},
			"with sort": {
				args: args{
					page:     1,
					pageSize: 10,
					sort:     "-id",
					query:    "",
				},
				want: &model.ListResult[model.Book]{
					Pagination: model.Pagination{
						Page:         1,
						PageSize:     10,
						TotalRecords: 2,
						TotalPages:   1,
						Offset:       0,
						Sort:         "id desc",
						HasNext:      false,
					},
					Data: []model.Book{
						{
							ID:     u.ID,
							Name:   "book2",
							Author: "author2",
							Topic: &model.Topic{
								ID:   topic.ID,
								Name: "topic1",
								Code: "code1",
							},
						},
						{
							ID:     u.ID - 1,
							Name:   "book1",
							Author: "author1",
							Topic: &model.Topic{
								ID:   topic.ID,
								Name: "topic1",
								Code: "code1",
							},
						},
					},
				},
				wantErr: false,
			},
			"with query": {
				args: args{
					page:     1,
					pageSize: 10,
					sort:     "-id",
					query:    "1",
				},
				want: &model.ListResult[model.Book]{
					Pagination: model.Pagination{
						Page:         1,
						PageSize:     10,
						TotalRecords: 1,
						TotalPages:   1,
						Offset:       0,
						Sort:         "id desc",
						HasNext:      false,
					},
					Data: []model.Book{
						{
							ID:     u.ID - 1,
							Name:   "book1",
							Author: "author1",
							Topic: &model.Topic{
								ID:   topic.ID,
								Name: "topic1",
								Code: "code1",
							},
						},
					},
				},
				wantErr: false,
			},
			"with page": {
				args: args{
					page:     2,
					pageSize: 1,
					sort:     "-id",
					query:    "",
				},
				want: &model.ListResult[model.Book]{
					Pagination: model.Pagination{
						Page:         2,
						PageSize:     1,
						TotalRecords: 2,
						TotalPages:   2,
						Offset:       1,
						Sort:         "id desc",
						HasNext:      false,
					},
					Data: []model.Book{
						{
							ID:     u.ID - 1,
							Name:   "book1",
							Author: "author1",
							Topic: &model.Topic{
								ID:   topic.ID,
								Name: "topic1",
								Code: "code1",
							},
						},
					},
				},
				wantErr: false,
			},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				r := &repo{}
				got, err := r.GetList(ctx, model.ListQuery{
					Page:     tt.args.page,
					Sort:     tt.args.sort,
					Query:    tt.args.query,
					PageSize: tt.args.pageSize,
				})
				if (err != nil) != tt.wantErr {
					t.Errorf("repo.GetList() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got.Pagination, tt.want.Pagination) {
					t.Errorf("%v case: repo.GetList() = %v, want %v", name, got.Pagination, tt.want.Pagination)
				}

				if !reflect.DeepEqual(len(got.Data), len(tt.want.Data)) {
					t.Errorf("repo.GetList() = %v, want %v", len(got.Data), len(tt.want.Data))
				}

				for i := 0; i < len(got.Data); i++ {
					if !reflect.DeepEqual(*got.Data[i].Topic, *tt.want.Data[i].Topic) {
						t.Errorf("repo.GetList() = %v, want %v at index", got.Data[i].Topic, tt.want.Data[i].Topic)
					}
					got.Data[i].Topic = nil
					tt.want.Data[i].Topic = nil

					if !reflect.DeepEqual(got.Data[i], tt.want.Data[i]) {
						t.Errorf("repo.GetList() = %v, want %v at index", got.Data[i], tt.want.Data[i])
					}
				}
			})
		}
	})
}

func Test_repo_IsExist(t *testing.T) {
	db.WithTestingDB(t, func(ctx db.Context) {
		topic := &orm.Topic{
			Name: "topic1",
			Code: "code1",
		}
		err := topic.Insert(ctx, ctx.DB, boil.Infer())
		require.NoError(t, err)

		u := &orm.Book{
			Name:    "book1",
			Author:  null.String{String: "author1", Valid: true},
			TopicID: null.Int{Int: topic.ID, Valid: true},
		}
		err = u.Insert(ctx, ctx.DB, boil.Infer())
		require.NoError(t, err)

		type args struct {
			uID int
		}
		tests := map[string]struct {
			args    args
			want    bool
			wantErr bool
		}{
			"success": {
				args: args{
					uID: u.ID,
				},
				want:    true,
				wantErr: false,
			},
			"not found": {
				args: args{
					uID: u.ID + 1,
				},
				want:    false,
				wantErr: false,
			},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				r := &repo{}
				got, err := r.IsExist(ctx, tt.args.uID)
				if (err != nil) != tt.wantErr {
					t.Errorf("repo.IsExist() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				require.Equal(t, tt.want, got)
			})
		}
	})
}

func Test_repo_Delete(t *testing.T) {
	db.WithTestingDB(t, func(ctx db.Context) {
		topic := &orm.Topic{
			Name: "topic1",
			Code: "code1",
		}
		err := topic.Insert(ctx, ctx.DB, boil.Infer())
		require.NoError(t, err)

		u := &orm.Book{
			Name:    "book1",
			Author:  null.String{String: "author1", Valid: true},
			TopicID: null.Int{Int: topic.ID, Valid: true},
		}
		err = u.Insert(ctx, ctx.DB, boil.Infer())
		require.NoError(t, err)

		type args struct {
			uID int
		}
		tests := map[string]struct {
			args    args
			want    bool
			wantErr bool
		}{
			"success": {
				args: args{
					uID: u.ID,
				},
				want:    true,
				wantErr: false,
			},
			"not found": {
				args: args{
					uID: u.ID + 1,
				},
				want:    false,
				wantErr: false,
			},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				r := &repo{}
				err := r.Delete(ctx, tt.args.uID)
				if (err != nil) != tt.wantErr {
					t.Errorf("repo.IsExist() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			})
		}
	})
}
