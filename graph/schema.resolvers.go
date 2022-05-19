package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/chloexu/hackernews/graph/generated"
	"github.com/chloexu/hackernews/graph/model"
	"github.com/rs/xid"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error) {
	n := len(r.Resolver.TodoStore)
	if n == 0 {
		r.Resolver.TodoStore = make(map[string]model.Todo)
	}

	var todo model.Todo
	nid := xid.New().String()
	todo.ID = nid
	todo.Text = input.Text
	todo.UserID = input.UserID
	todo.Done = false
	if todo.Done {
		todo.Done = true
	}
	r.Resolver.TodoStore[nid] = todo
	return &todo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error) {
	fmt.Sprintln("enter UpsertTodo")
	id := input.ID
	var todo model.Todo

	n := len(r.Resolver.TodoStore)
	if n == 0 {
		r.Resolver.TodoStore = make(map[string]model.Todo)
	}

	todo, ok := r.Resolver.TodoStore[id]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	if input.Text != nil {
		todo.Text = *input.Text
	}
	if input.Done != nil {
		todo.Done = *input.Done
	}
	r.Resolver.TodoStore[id] = todo
	return &todo, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	todo, ok := r.Resolver.TodoStore[id]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return &todo, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }