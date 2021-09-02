package repos

import (
	"fmt"
	"github.com/asavt7/todo/pkg/domain"
	"github.com/jmoiron/sqlx"
)

type TodoRepoPostgres struct {
	db *sqlx.DB
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
