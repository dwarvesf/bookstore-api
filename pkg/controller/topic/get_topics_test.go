package topic

import (
	"errors"
	"net/http/httptest"
	"testing"

	mocks "github.com/dwarvesf/bookstore-api/mocks/pkg/repository/topic"
	"github.com/dwarvesf/bookstore-api/pkg/config"
	"github.com/dwarvesf/bookstore-api/pkg/handler/testutil"
	"github.com/dwarvesf/bookstore-api/pkg/logger/monitor"
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/dwarvesf/bookstore-api/pkg/repository"
	"github.com/dwarvesf/bookstore-api/pkg/repository/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_impl_GetTopics(t *testing.T) {
	type mocked struct {
		expGetTopicsCalled bool
		getTopics          []model.Topic
		getTopicsErr       error
	}

	tests := map[string]struct {
		mocked  mocked
		wantErr bool
	}{
		"success": {
			mocked: mocked{
				expGetTopicsCalled: true,
				getTopics: []model.Topic{
					{
						ID:   1,
						Name: "topic1",
						Code: "code1",
					},
					{
						ID:   2,
						Name: "topic2",
						Code: "code2",
					},
				},
			},
			wantErr: false,
		},
		"failed": {
			mocked: mocked{
				expGetTopicsCalled: true,
				getTopicsErr:       errors.New("error"),
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var (
				topicRepoMock = mocks.NewRepo(t)
			)

			if tt.mocked.expGetTopicsCalled {
				topicRepoMock.
					EXPECT().
					GetAll(mock.Anything).
					Return(tt.mocked.getTopics, tt.mocked.getTopicsErr)
			}

			c := &impl{
				repo: &repository.Repo{
					Topic: topicRepoMock,
				},
				cfg:     config.LoadTestConfig(),
				monitor: monitor.TestMonitor(),
			}

			_, err := db.Init(c.cfg)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			ginCtx := testutil.NewRequest(w, testutil.MethodPost, nil, nil, nil, nil)
			books, err := c.GetTopics(ginCtx.Request.Context())
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.GetTopics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.mocked.getTopics, books)
		})
	}
}
