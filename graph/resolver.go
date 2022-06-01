package graph

import (
	"github.com/chloexu/hackernews/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	// TodoStore map[string]model.Todo
	Repo repository.Repository
}
