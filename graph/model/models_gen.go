// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateTodoInput struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
	Done   *bool  `json:"done"`
}

type Todo struct {
	ID          string `json:"id"`
	Text        string `json:"text"`
	Done        bool   `json:"done"`
	UserID      string `json:"userId"`
	CreatedAt   string `json:"createdAt"`
	CompletedAt string `json:"completedAt"`
}

type UpdateTodoInput struct {
	ID   string  `json:"id"`
	Text *string `json:"text"`
	Done *bool   `json:"done"`
}
