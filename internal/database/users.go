package database

import (
	"context"
	"database/sql"
	"time"
)

type UserModel struct {
	DB *sql.DB
}

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"`
}

func (m *UserModel) Insert(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := `INSERT INTO users (email, name, password) VALUES (?, ?, ?) RETURNING id`

	return m.DB.QueryRowContext(ctx, query, user.Email, user.Name, user.Password).Scan(&user.Id)

}

func (m *UserModel) Get(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM users WHERE Id = ?`

	var user User

	err := m.DB.QueryRowContext(ctx, query, id).Scan(&user.Id, &user.Email, &user.Name, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
