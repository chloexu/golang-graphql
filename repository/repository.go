package repository

import "time"

type TodoRow struct {
	ID          string
	Text        string
	Done        bool
	UserID      string
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Repository interface {
	TodoByID(id string) (TodoRow, error)
	TodosByUser(userId string) ([]TodoRow, error)
	AddTodo(row TodoRow) (bool, error)
	UpdateTodo(row TodoRow) (bool, error)
	Close()
}
