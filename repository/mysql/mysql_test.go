package mysql

import (
	"database/sql"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	repo "github.com/chloexu/hackernews/repository"
)

var createdAt time.Time
var completedAt time.Time
var todo = &repo.TodoRow{
	ID:          "caajol287d5nser73bs0",
	UserID:      "chloexu1124",
	Text:        "Water roses and lilies",
	CreatedAt:   createdAt,
	CompletedAt: completedAt,
	Done:        false,
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error %s was not expected when opening a stub database", err)
	}

	return db, mock
}

func TestTodoByID(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		id string
	}
	db, mock := NewMock()
	repos := &mysqlRepository{db}

	defer func() {
		repos.Close()
	}()

	query := "SELECT id, text, done, user_id, created_at, completed_at FROM todos WHERE id = \\?"
	rows := sqlmock.NewRows([]string{"id", "text", "done", "userId", "created_at", "completed_at"}).
		AddRow(todo.ID, todo.Text, todo.Done, todo.UserID, todo.CreatedAt, todo.CompletedAt)
	mock.ExpectQuery(query).WithArgs(todo.ID).WillReturnRows(rows)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    repo.TodoRow
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test todo by id", fields{db}, args{id: "caajol287d5nser73bs0"}, *todo, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mysqlRepository{
				db: tt.fields.db,
			}
			got, err := r.TodoByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.TodoByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.TodoByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
