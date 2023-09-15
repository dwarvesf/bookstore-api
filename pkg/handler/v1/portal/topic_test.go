package portal

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	mocks "github.com/dwarvesf/bookstore-api/mocks/pkg/controller/topic"
	"github.com/dwarvesf/bookstore-api/pkg/config"
	"github.com/dwarvesf/bookstore-api/pkg/handler/testutil"
	"github.com/dwarvesf/bookstore-api/pkg/handler/v1/view"
	"github.com/dwarvesf/bookstore-api/pkg/logger"
	"github.com/dwarvesf/bookstore-api/pkg/logger/monitor"
	"github.com/dwarvesf/bookstore-api/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_GetTopics(t *testing.T) {
	type mocked struct {
		expGetTopics bool
		topics       []*model.Topic
		getTopicsErr error
	}

	type expected struct {
		Status  int
		Body    view.GetTopicsResponse
		WantErr bool
		Err     string
	}
	tests := map[string]struct {
		mocked   mocked
		expected expected
	}{
		"success": {
			mocked: mocked{
				expGetTopics: true,
				topics: []*model.Topic{
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
			expected: expected{
				Status: 200,
				Body: view.GetTopicsResponse{
					Data: []*view.Topic{
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
			},
		},
		// "failed to get": {
		// 	mocked: mocked{
		// 		page:        "1",
		// 		pageSize:    "10",
		// 		sort:        "created_at desc",
		// 		query:       "",
		// 		expGetTopics: true,
		// 		getTopicsErr: errors.New("failed to get books"),
		// 	},
		// 	expected: expected{
		// 		Status:  500,
		// 		WantErr: true,
		// 		Err:     "INTERNAL_ERROR",
		// 	},
		// },
	}
	for name, tt := range tests {
		w := httptest.NewRecorder()
		cfg := config.LoadTestConfig()

		ginCtx := testutil.NewRequest(w, testutil.MethodGet, nil, nil, nil, nil)

		var (
			ctrlMock = mocks.NewController(t)
		)

		if tt.mocked.expGetTopics {
			ctrlMock.EXPECT().GetTopics(mock.Anything).Return(tt.mocked.topics, tt.mocked.getTopicsErr)
		}
		t.Run(name, func(t *testing.T) {
			h := Handler{
				log:       logger.NewLogger(),
				cfg:       cfg,
				topicCtrl: ctrlMock,
				monitor:   monitor.TestMonitor(),
			}
			h.GetTopics(ginCtx)

			assert.Equal(t, tt.expected.Status, w.Code)
			resBody := w.Body.String()
			body, err := json.Marshal(tt.expected.Body)
			assert.Nil(t, err)
			if !tt.expected.WantErr {
				assert.Equal(t, string(body), resBody)
			} else {
				assert.Contains(t, resBody, tt.expected.Err)
			}
		})
	}
}
