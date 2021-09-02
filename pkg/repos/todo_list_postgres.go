package repos

import (
	"fmt"
	"github.com/asavt7/todo/pkg/domain"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func (t *TodoListPostgres) Update(userId int, listId int, input domain.UpdateTodoListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s l SET %s FROM %s ul WHERE l.id=ul.list_id AND ul.user_id=$%d AND ul.list_id=$%d", todoListsTable, setQuery, userListsTable, argId, argId+1)

	args = append(args, userId, listId )
	logrus.Debugf("todo list update sql: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := t.db.Exec(query, args...)
	return err
}

func (t *TodoListPostgres) Delete(userId int, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s l USING %s ul WHERE l.id = ul.list_id AND ul.user_id=$1 AND ul.id=$2", todoListsTable, userListsTable)
	result, err := t.db.Exec(query, listId, userId)

	logrus.Info(result)
	return err
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
