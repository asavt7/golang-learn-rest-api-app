package domain

import "errors"

type ID int

type TodoList struct {
	Id          ID     `json:"id" db:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" db:"description"`
}
type UpdateTodoListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (u *UpdateTodoListInput) Validate() error {
	if u.Title == nil && u.Description == nil {
		return errors.New("update input has no values")
	}
	return nil
}

type UserList struct {
	Id         ID `db:"id"`
	UserId     ID `db:"user_id"`
	TodoListId ID `db:"list_id"`
}

type TodoItem struct {
	Id          ID     `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required" test:"yesTest"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type UpdateTodoItem struct {
	Title       *string `json:"title" `
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (u *UpdateTodoItem) Validate() error {
	if u.Title == nil && u.Description == nil && u.Done == nil {
		return errors.New("update input has no values")
	}
	return nil
}

type TodoListItem struct {
	Id         ID
	TodoListId ID
	ItemId     ID
}
