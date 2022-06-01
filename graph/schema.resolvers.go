package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/chloexu/hackernews/graph/generated"
	"github.com/chloexu/hackernews/graph/model"
	"github.com/chloexu/hackernews/repository"
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

	var row repository.TodoRow
	nid := xid.New().String()
	row.ID = nid
	row.Text = input.Text
	row.UserID = input.UserID
	row.Done = false
	row.CreatedAt = time.Now()
	row.CompletedAt = time.Now()
	// isSuccessful, err := data.AddTodo(row)
	isSuccessful, err := r.Repo.AddTodo(row)
	if err != nil {
		return nil, fmt.Errorf("CreateTodo failed %v", err)
	}
	if !isSuccessful {
		return nil, fmt.Errorf("CreateTodo no record inserted")
	}
	inserted, err := r.Repo.TodoByID(nid)
	if err != nil {
		return nil, fmt.Errorf("CreateTodo failed to get todo %q %v", nid, err)
	}
	todo := &model.Todo{
		ID:          inserted.ID,
		Text:        inserted.Text,
		UserID:      inserted.UserID,
		Done:        inserted.Done,
		CreatedAt:   inserted.CreatedAt.Format("2006-01-02 15:04:05"),
		CompletedAt: inserted.CompletedAt.Format("2006-01-02 15:04:05"),
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
	var row repository.TodoRow
	row.ID = input.ID
	if input.Text != nil {
		row.Text = *input.Text
	} else {
		row.Text = ""
	}
	row.Done = input.Done
	// isSuccessful, err := data.UpdateTodo(input)
	isSuccessful, err := r.Repo.UpdateTodo(row)
	if err != nil {
		return nil, fmt.Errorf("UpdateTodo failed to update todo %q, %v", input.ID, err)
	}
	if !isSuccessful {
		return nil, fmt.Errorf("UpdateTodo no record to update")
	}
	row, err = r.Repo.TodoByID(input.ID)
	if err != nil {
		return nil, fmt.Errorf("UpdateTodo failed to get todo %q, %v", input.ID, err)
	}
	todo := &model.Todo{
		ID:          row.ID,
		Text:        row.Text,
		UserID:      row.UserID,
		Done:        row.Done,
		CreatedAt:   row.CreatedAt.Format("2006-01-02 15:04:05"),
		CompletedAt: row.CompletedAt.Format("2006-01-02 15:04:05"),
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
	// row, err := data.TodoByID(id)
	row, err := r.Repo.TodoByID(id)
	if err != nil {
		return nil, fmt.Errorf("Todo Failed to retrieve TodoByID %q, %v", id, err)
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
	// todoRows, err := data.TodosByUser(userID)
	todoRows, err := r.Repo.TodosByUser(userID)
	if err != nil {
		return nil, fmt.Errorf("Todos Failed to retrieve todos: %v", err)
	}
	todos := make([]*model.Todo, 0)
	for _, row := range todoRows {
		todo := &model.Todo{
			ID:          row.ID,
			Text:        row.Text,
			UserID:      row.UserID,
			Done:        row.Done,
			CreatedAt:   row.CreatedAt.Format("2006-01-02 15:04:05"),
			CompletedAt: row.CompletedAt.Format("2006-01-02 15:04:05"),
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
