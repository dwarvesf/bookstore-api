package view

// Topic represent the topic
type Topic struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
} // @name Topic

// GetTopicsResponse represent the get topics response
type GetTopicsResponse struct {
	Data []*Topic `json:"data"`
} // @name GetTopicsResponse
