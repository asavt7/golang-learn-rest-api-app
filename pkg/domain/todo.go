package domain

type ID int

type TodoList struct {
	Id          ID     `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type UserList struct {
	Id         ID
	UserId     ID
	TodoListId ID
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
