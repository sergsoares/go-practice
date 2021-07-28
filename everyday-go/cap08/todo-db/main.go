package main

import (
	"database/sql"
	"flag"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Todo struct {
	ID            int
	Description   string
	CreateDate    time.Time
	CompletedDate *time.Time
}

func main() {
	db := connect()

	var todo string
	var action string

	flag.StringVar(&todo, "todo", "", "")
	flag.StringVar(&action, "action", "", "")
	flag.Parse()

	switch action {
	case "create":
		create(db, todo)
	case "list":
		todo, _ := list(db)
		for _, k := range todo {
			fmt.Println("-", k.Description, k.CreateDate.String())
		}
	default:
		panic("Supported actions are: list, create")
	}

	// log.Debug().Msg("Fibonacci")
}

func list(db *sql.DB) ([]Todo, error) {
	rows, err := db.Query("SELECT id, description, created_date, completed_date FROM todo")
	if err != nil {
		return []Todo{}, fmt.Errorf("Unable to get from todo table: %w", err)
	}

	todos := []Todo{}
	defer rows.Close()
	for rows.Next() {
		result := Todo{}

		err := rows.Scan(
			&result.ID,
			&result.Description,
			&result.CreateDate,
			&result.CompletedDate,
		)

		if err != nil {
			return []Todo{}, fmt.Errorf("Row scan error: %w", err)
		}

		todos = append(todos, result)
	}

	return todos, nil
}

func create(db *sql.DB, todo string) error {
	res, err := db.Query(`Insert into
	todo
	(id, description, created_date)
	VALUES
	(DEFAULT, $1, now());`,
		todo)

	if err != nil {
		return err
	}
	defer res.Close()

	return nil
}

func connect() *sql.DB {
	url := "postgresql://postgres:postgres@localhost:5432/db?sslmode=disable"

	db, _ := sql.Open("postgres", url)
	err := db.Ping()

	if err != nil {
		panic(err)
	}

	// log.Debug().Msg("ok")
	return db
}
