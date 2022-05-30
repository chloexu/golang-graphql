package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/chloexu/hackernews/data"
	"github.com/chloexu/hackernews/graph/generated"
	"github.com/chloexu/hackernews/graph/model"
	"github.com/rs/xid"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error) {
	// n := len(r.Resolver.TodoStore)
	// if n == 0 {
	// 	r.Resolver.TodoStore = make(map[string]model.Todo)
	// }

	// var todo model.Todo
	// nid := xid.New().String()
	// todo.ID = nid
	// todo.Text = input.Text
	// todo.UserID = input.UserID
	// todo.Done = false
	// currentTime := time.Now()
	// todo.CreatedAt = currentTime.Format("2006-01-02 15:04:05")
	// if todo.Done {
	// 	todo.Done = true
	// }
	// r.Resolver.TodoStore[nid] = todo
	// return &todo, nil

	var row data.TodoRow
	nid := xid.New().String()
	row.ID = nid
	row.Text = input.Text
	row.UserID = input.UserID
	row.Done = false
	row.CreatedAt = time.Now()
	row, err := data.AddTodo(row)
	if err != nil {
		return nil, fmt.Errorf("CreateTodo %v", err)
	}
	todo := &model.Todo{
		ID:          row.ID,
		Text:        row.Text,
		UserID:      row.UserID,
		Done:        row.Done,
		CreatedAt:   row.CreatedAt.Format("2006-01-02 15:04:05"),
		CompletedAt: "",
	}
	return todo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error) {
	// fmt.Sprintln("enter UpsertTodo")
	// id := input.ID
	// var todo model.Todo

	// n := len(r.Resolver.TodoStore)
	// if n == 0 {
	// 	r.Resolver.TodoStore = make(map[string]model.Todo)
	// }

	// todo, ok := r.Resolver.TodoStore[id]
	// if !ok {
	// 	return nil, fmt.Errorf("not found")
	// }
	// if input.Text != nil {
	// 	todo.Text = *input.Text
	// }
	// if input.Done != nil {
	// 	todo.Done = *input.Done
	// 	if *input.Done == true {
	// 		currentTime := time.Now()
	// 		todo.CompletedAt = currentTime.Format("2006-01-02 15:04:05")
	// 	} else {
	// 		todo.CompletedAt = ""
	// 	}
	// }
	// r.Resolver.TodoStore[id] = todo
	// return &todo, nil

	row, err := data.UpdateTodo(input)
	if err != nil {
		return nil, fmt.Errorf("Failed to update todo %q %v", input.ID, err)
	}
	todo := &model.Todo{
		ID:          row.ID,
		Text:        row.Text,
		UserID:      row.UserID,
		Done:        row.Done,
		CreatedAt:   row.CreatedAt.Format("2006-01-02 15:04:05"),
		CompletedAt: "",
	}

	if !row.CompletedAt.Time.IsZero() {
		todo.CompletedAt = row.CompletedAt.Time.Format("2006-01-02 15:04:05")
	}

	return todo, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	// START - USING IN-MEMORY STORE
	// todo, ok := r.Resolver.TodoStore[id]
	// if !ok {
	// 	return nil, fmt.Errorf("not found")
	// }
	// return &todo, nil
	// END - USING IN-MEMORY STORE

	// START - USING LOCAL DB
	row, err := data.TodoByID(id)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve TodoByID %q %v", id, err)
	}
	todo := &model.Todo{ // ???
		ID:          row.ID,
		Text:        row.Text,
		UserID:      row.UserID,
		Done:        row.Done,
		CreatedAt:   row.CreatedAt.Format("2006-01-02 15:04:05"),
		CompletedAt: "",
	}
	return todo, nil
	// END - USING LOCAL DB
}

func (r *queryResolver) Todos(ctx context.Context, userID string) ([]*model.Todo, error) {
	// START - USING IN-MEMORY STORE
	// n := len(r.Resolver.TodoStore)
	// if n == 0 {
	// 	r.Resolver.TodoStore = make(map[string]model.Todo)
	// }
	// todos := make([]*model.Todo, 0)
	// for id := range r.Resolver.TodoStore {
	// 	todo, ok := r.Resolver.TodoStore[id]
	// 	if !ok {
	// 		return nil, fmt.Errorf("not found")
	// 	}
	// 	if todo.UserID == userID {
	// 		todos = append(todos, &todo)
	// 	}
	// }
	// END - USING IN-MEMORY STORE

	// START - USING LOCAL DB
	todoRows, err := data.TodosByUser(userID)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve todos: %v", err)
	}
	todos := make([]*model.Todo, 0)
	for _, row := range todoRows {
		todo := &model.Todo{
			ID:          row.ID,
			Text:        row.Text,
			UserID:      row.UserID,
			Done:        row.Done,
			CreatedAt:   row.CreatedAt.Format("2006-01-02 15:04:05"),
			CompletedAt: "",
		}
		if !row.CompletedAt.Time.IsZero() {
			todo.CompletedAt = row.CompletedAt.Time.Format("2006-01-02 15:04:05")
		}
		todos = append(todos, todo)
	}
	return todos, nil
	// END - USING LOCAL DB
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
