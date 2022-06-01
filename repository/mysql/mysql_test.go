package mysql

import (
	"database/sql"
	"fmt"
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
var todoBySameUser = &repo.TodoRow{
	ID:          "caajol287d5nser73fh9",
	UserID:      "chloexu1124",
	Text:        "Pick up laundry",
	CreatedAt:   createdAt,
	CompletedAt: completedAt,
	Done:        false,
}
var todoByDifferentUser = &repo.TodoRow{
	ID:          "caajol287d5nsergf35",
	UserID:      "1124chloezhuqing",
	Text:        "Water roses and lilies",
	CreatedAt:   createdAt,
	CompletedAt: completedAt,
	Done:        false,
}
var todoUpdateTextDone = &repo.TodoRow{
	ID:   "caajol287d5nser73bs0",
	Text: "Pick up laundry",
	Done: true,
}
var todoUpdateTextNotDone = &repo.TodoRow{
	ID:   "caajol287d5nser73bs0",
	Text: "Pick up laundry 2",
	Done: false,
}
var todoUpdateDone = &repo.TodoRow{
	ID:   "caajol287d5nser73bs0",
	Done: true,
}
var todoUpdateNotDone = &repo.TodoRow{
	ID:   "caajol287d5nser73bs0",
	Done: false,
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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
	mysqlRepo := &mysqlRepository{db}

	defer func() {
		mysqlRepo.Close()
	}()

	query := "SELECT id, text, done, user_id, created_at, completed_at FROM todos WHERE id = ?"
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

func TestTodosByUser(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		userId string
	}

	db, mock := NewMock()
	mysqlRepo := &mysqlRepository{db}

	defer func() {
		mysqlRepo.Close()
	}()

	query := "SELECT  id, text, done, user_id, created_at, completed_at FROM todos WHERE user_id = ?"

	rows := sqlmock.NewRows([]string{"id", "text", "done", "userId", "created_at", "completed_at"}).
		AddRow(todo.ID, todo.Text, todo.Done, todo.UserID, todo.CreatedAt, todo.CompletedAt).
		AddRow(todoBySameUser.ID, todoBySameUser.Text, todoBySameUser.Done, todoBySameUser.UserID,
			todoBySameUser.CreatedAt, todoBySameUser.CompletedAt)
	mock.ExpectQuery(query).WithArgs(todo.UserID).WillReturnRows(rows)

	rowsOfDiffUser := sqlmock.NewRows([]string{"id", "text", "done", "userId", "created_at", "completed_at"}).
		AddRow(todoByDifferentUser.ID, todoByDifferentUser.Text, todoByDifferentUser.Done, todoByDifferentUser.UserID,
			todoBySameUser.CreatedAt, todoBySameUser.CompletedAt)
	mock.ExpectQuery(query).WithArgs(todoByDifferentUser.UserID).WillReturnRows(rowsOfDiffUser)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []repo.TodoRow
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test todo by user should return 2 rows",
			fields{db},
			args{userId: todo.UserID},
			[]repo.TodoRow{*todo, *todoBySameUser},
			false,
		},
		{"test todo by user should return 1 row",
			fields{db},
			args{userId: todoByDifferentUser.UserID},
			[]repo.TodoRow{*todoByDifferentUser},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mysqlRepository{
				db: tt.fields.db,
			}
			got, err := r.TodosByUser(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlRepository.TodosByUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlRepository.TodosByUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddTodo(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		row repo.TodoRow
	}

	db, mock := NewMock()
	mysqlRepo := &mysqlRepository{db}

	defer func() {
		mysqlRepo.Close()
	}()

	statement := "INSERT INTO todos(id, text, done, user_id, created_at, completed_at) VALUES (?, ?, ?, ?, curdate(), curdate())"

	mock.ExpectExec(statement).WithArgs(
		todo.ID, todo.Text, todo.Done, todo.UserID,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{"test add todo by user should success",
			fields{db},
			args{*todo},
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mysqlRepository{
				db: tt.fields.db,
			}
			fmt.Println(tt.args.row)
			got, err := r.AddTodo(tt.args.row)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlRepository.AddTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("mysqlRepository.AddTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateTodo(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		row repo.TodoRow
	}

	db, mock := NewMock()
	mysqlRepo := &mysqlRepository{db}

	defer func() {
		mysqlRepo.Close()
	}()

	statement1 := "UPDATE todos SET text = ?, done = ?, completed_at = curdate() where id = ?"
	mock.ExpectExec(statement1).WithArgs(todoUpdateTextDone.Text, todoUpdateTextDone.Done, todoUpdateTextDone.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	statement2 := "UPDATE todos SET text = ?, done = ?, completed_at = null where id = ?"
	mock.ExpectExec(statement2).WithArgs(todoUpdateTextNotDone.Text, todoUpdateTextNotDone.Done, todoUpdateTextNotDone.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	statement3 := "UPDATE todos SET done = ?, completed_at = curdate() where id = ?"
	mock.ExpectExec(statement3).WithArgs(todoUpdateDone.Done, todoUpdateDone.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	statement4 := "UPDATE todos SET done = ?, completed_at = null where id = ?"
	mock.ExpectExec(statement4).WithArgs(todoUpdateNotDone.Done, todoUpdateNotDone.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			"test update todo text and done to true: should return success",
			fields{db},
			args{*todoUpdateTextDone},
			true,
			false,
		},
		{
			"test update todo text and done to false: should return success",
			fields{db},
			args{*todoUpdateTextNotDone},
			true,
			false,
		},
		{
			"test update todo done to true: should return success",
			fields{db},
			args{*todoUpdateDone},
			true,
			false,
		},
		{
			"test update todo done to false: should return success",
			fields{db},
			args{*todoUpdateNotDone},
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mysqlRepository{
				db: tt.fields.db,
			}
			got, err := r.UpdateTodo(tt.args.row)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlRepository.UpdateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("mysqlRepository.UpdateTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}
