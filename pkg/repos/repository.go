package repos

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repo struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepo() *Repo {
	return new(Repo)
}
