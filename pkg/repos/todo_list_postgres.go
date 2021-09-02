package repos

import (
	"fmt"
	"github.com/asavt7/todo/pkg/domain"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func (t *TodoListPostgres) GetListById(userId int, listId int) (domain.TodoList, error) {
	var list domain.TodoList
	query := fmt.Sprintf("SELECT l.id, l.title, l.description FROM %s l INNER JOIN %s ul ON l.id = ul.list_id WHERE ul.user_id=$1 AND ul.id=$2", todoListsTable, userListsTable)
	err := t.db.Get(&list, query, userId, listId)
	return list, err
}

func (t *TodoListPostgres) GetAllLists(userId int) ([]domain.TodoList, error) {

	var lists []domain.TodoList
	query := fmt.Sprintf("SELECT l.id, l.title, l.description FROM %s l INNER JOIN %s ul ON l.id = ul.list_id WHERE ul.user_id=$1", todoListsTable, userListsTable)
	err := t.db.Select(&lists, query, userId)
	if err != nil {
		return nil, err
	}
	return lists, nil
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (t *TodoListPostgres) Create(userId int, input domain.TodoList) (int, error) {

	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1,$2) RETURNING ID", todoListsTable)
	row := tx.QueryRow(createListQuery, input.Title, input.Description)

	var listId int
	err = row.Scan(&listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createUserListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", userListsTable)
	_, err = tx.Exec(createUserListQuery, userId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return listId, nil
}
