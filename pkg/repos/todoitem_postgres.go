package repos

import (
	"fmt"
	"github.com/asavt7/todo/pkg/domain"
	"github.com/jmoiron/sqlx"
)

type TodoRepoPostgres struct {
	db *sqlx.DB
}

func (t *TodoRepoPostgres) Delete(userId int, itemId int) error {

	query := fmt.Sprintf("DELETE FROM %s ti USING %s li , %s ul WHERE ti.id=li.item_id AND ul.list_id=li.list_id AND  ul.user_id=$1 AND ti.id=$2", todoItemsTable, listsItemTable, userListsTable)

	_, err := t.db.Exec(query, userId, itemId)
	return err
}

func (t *TodoRepoPostgres) GetById(userId int, itemId int) (domain.TodoItem, error) {
	query := fmt.Sprintf("SELECT ti.* FROM %s ti INNER JOIN %s li ON ti.id=li.item_id  INNER JOIN %s ul ON ul.list_id=li.list_id WHERE ul.user_id=$1 AND ti.id=$2", todoItemsTable, listsItemTable, userListsTable)
	var result domain.TodoItem
	err := t.db.Get(&result, query, userId, itemId)
	return result, err
}

func (t *TodoRepoPostgres) GetAllItems(userId int, listId int) ([]domain.TodoItem, error) {
	query := fmt.Sprintf("SELECT ti.* FROM %s ti INNER JOIN %s li ON ti.id=li.item_id  INNER JOIN %s ul ON ul.list_id=li.list_id WHERE ul.user_id=$1 AND li.list_id=$2", todoItemsTable, listsItemTable, userListsTable)
	var result []domain.TodoItem
	err := t.db.Select(&result, query, userId, listId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t *TodoRepoPostgres) Create(listId int, input domain.TodoItem) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING ID", todoItemsTable)
	row := tx.QueryRow(createItemQuery, input.Title, input.Description)
	var idItem int
	err = row.Scan(&idItem)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	insertItemListQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listsItemTable)
	_, err = tx.Exec(insertItemListQuery, listId, idItem)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return idItem, tx.Commit()
}

func NewTodoRepoPostgres(db *sqlx.DB) *TodoRepoPostgres {
	return &TodoRepoPostgres{db: db}
}
