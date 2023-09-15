package topic

import (
	"testing"

	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/repository/db"
	"github.com/dwarvesf/bookstore-api/pkg/repository/orm"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Test_repo_IsExist(t *testing.T) {
	db.WithTestingDB(t, func(ctx db.Context) {
		topic := &orm.Topic{
			Name: "topic1",
			Code: "code1",
		}
		err := topic.Insert(ctx, ctx.DB, boil.Infer())
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
					uID: topic.ID,
				},
				want:    true,
				wantErr: false,
			},
			"not found": {
				args: args{
					uID: topic.ID + 1,
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

func Test_repo_GetAll(t *testing.T) {
	db.WithTestingDB(t, func(ctx db.Context) {
		tests := map[string]struct {
			want    []*model.Topic
			wantErr bool
		}{
			"success": {
				want: []*model.Topic{
					{
						ID:   1,
						Name: "Programming",
						Code: "programming",
					},
					{
						ID:   2,
						Name: "Database",
						Code: "database",
					},
					{
						ID:   3,
						Name: "DevOps",
						Code: "devops",
					},
				},
				wantErr: false,
			},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				r := &repo{}
				got, err := r.GetAll(ctx)
				if (err != nil) != tt.wantErr {
					t.Errorf("repo.GetAll() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				require.Equal(t, tt.want, got)
			})
		}
	})
}
