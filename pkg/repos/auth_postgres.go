package repos

import (
	"fmt"
	"github.com/asavt7/todo/pkg/domain"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a AuthPostgres) CreateUser(user domain.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username,password_hash) values( $1,$2, $3) RETURNING id ", usersTable)
	row := a.db.QueryRow(query, user.Name, user.Username, user.Password)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
