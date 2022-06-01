package data

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/chloexu/hackernews/graph/model"
// 	"github.com/go-sql-driver/mysql"
// )

// var db *sql.DB

// type TodoRow struct {
// 	ID          string
// 	Text        string
// 	Done        bool
// 	UserID      string
// 	CreatedAt   time.Time
// 	CompletedAt sql.NullTime
// }

// func TodoByID(id string) (TodoRow, error) {
// 	var todo TodoRow
// 	row := db.QueryRow("SELECT id, text, done, user_id, created_at, completed_at FROM todos WHERE id = ?", id)
// 	if err := row.Scan(&todo.ID, &todo.Text, &todo.Done, &todo.UserID, &todo.CreatedAt, &todo.CompletedAt); err != nil {
// 		if err == sql.ErrNoRows {
// 			return todo, fmt.Errorf("TodoByID row scan: no row. %q %v", id, err)
// 		}
// 		return todo, fmt.Errorf("TodoByID row scan: %q %v", id, err)
// 	}
// 	return todo, nil
// }

// func TodosByUser(userId string) ([]TodoRow, error) {
// 	// define todos slice to hold data from returned rows
// 	var todos []TodoRow

// 	/// read data from db
// 	rows, err := db.Query("SELECT id, text, done, user_id, created_at, completed_at FROM todos WHERE user_id = ?", userId)
// 	if err != nil {
// 		return nil, fmt.Errorf("TodosByUsers query %q: %v", userId, err)
// 	}

// 	defer rows.Close()

// 	// loop through rows, using Scan to assign column data to struct fields
// 	for rows.Next() {
// 		var todo TodoRow
// 		if err := rows.Scan(&todo.ID, &todo.Text, &todo.Done, &todo.UserID, &todo.CreatedAt, &todo.CompletedAt); err != nil {
// 			return nil, fmt.Errorf("TodosByUsers scan row %q: %v", userId, err)
// 		}
// 		todos = append(todos, todo)
// 	}
// 	// After the loop, check for an error from the overall query, using rows.Err.
// 	// Note that if the query itself fails, checking for an error here is the only
// 	// way to find out that the results are incomplete.
// 	if err := rows.Err(); err != nil {
// 		return nil, fmt.Errorf("TodosByUsers rows err %q: %v", userId, err)
// 	}

// 	return todos, nil
// }

// func AddTodo(row TodoRow) (bool, error) {
// 	result, err := db.Exec("INSERT INTO todos(id, text, done, user_id, created_at) VALUES (?, ?, ?, ?, ?)", row.ID, row.Text,
// 		row.Done, row.UserID, time.Now())

// 	if err != nil {
// 		return false, fmt.Errorf("AddTodo exec : %v", err)
// 	}

// 	inserted, err := result.RowsAffected()
// 	if err != nil {
// 		return false, fmt.Errorf("AddTodo fetch row after insertion : %v", err)
// 	}
// 	if inserted > 0 {
// 		return true, nil
// 	}
// 	return false, nil
// }

// func UpdateTodo(input model.UpdateTodoInput) (bool, error) {
// 	var result sql.Result
// 	if input.Done {
// 		r1, err := db.Exec("UPDATE todos SET text = ?, done = ?, completed_at = ? where id = ?", input.Text, input.Done, time.Now(), input.ID)
// 		if err != nil {
// 			return false, fmt.Errorf("UpdateTodo exec : %v", err)
// 		}
// 		result = r1
// 	} else {
// 		r2, err := db.Exec("UPDATE todos SET text = ?, done = ?, completed_at = null where id = ?", input.Text, input.Done, input.ID)
// 		if err != nil {
// 			return false, fmt.Errorf("UpdateTodo exec : %v", err)
// 		}
// 		result = r2
// 	}

// 	updated, err := result.RowsAffected()
// 	if err != nil {
// 		return false, fmt.Errorf("UpdateTodo fetch row after update : %v", err)
// 	}
// 	if updated > 0 {
// 		return true, nil
// 	}
// 	return false, nil
// }

// func InitializeDB() {
// 	// Capture connection properties
// 	cfg := mysql.Config{
// 		User:      os.Getenv("DBUSER"),
// 		Passwd:    os.Getenv("DBPASS"),
// 		Net:       "tcp",
// 		Addr:      "127.0.0.1:3306",
// 		DBName:    "todos_db",
// 		ParseTime: true,
// 	}

// 	// Get a database handle.
// 	var err error
// 	db, err = sql.Open("mysql", cfg.FormatDSN())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	pingErr := db.Ping()
// 	if pingErr != nil {
// 		log.Fatal(pingErr)
// 	}
// 	log.Println("DB connection established.")
// }
