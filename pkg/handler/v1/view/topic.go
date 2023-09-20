package view

// Topic represent the topic
type Topic struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
} // @name Topic

// TopicsResponse is the response for get topics
type TopicsResponse = Response[[]Topic] // @name TopicsResponse
