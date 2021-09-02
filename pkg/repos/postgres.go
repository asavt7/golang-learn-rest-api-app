package repos

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	usersTable     = "users"
	todoListsTable = "todo_lists"
	todoItemsTable = "todo_items"
	listsItemTable = "lists_item"
	userListsTable = "user_lists"
)

type Config struct {
	Host, Port, Username, Password, DBName string
	SSLMode                                string
}

func NewPostgreDb(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewPostgresRepo(db *sqlx.DB) *Repo {
	return &Repo{
		Authorization: NewAuthPostgres(db),
		TodoList : NewTodoListPostgres(db),
	}
}


