package domain

type ID int

type TodoList struct {
	Id          ID     `json:"id" db:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" db:"description"`
}

type UserList struct {
	Id         ID `db:"id"`
	UserId     ID `db:"user_id"`
	TodoListId ID `db:"list_id"`
}

type TodoItem struct {
	Id          ID     `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool
}

type TodoListItem struct {
	Id         ID
	TodoListId ID
	ItemId     ID
}
