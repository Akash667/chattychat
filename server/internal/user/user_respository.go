package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {

	return &repository{db: db}

}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {

	query := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, user.Username, user.Password, user.Email).Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (r *repository) GetUser(ctx context.Context, email string) (*User, error) {

	query := `SELECT id, username, email, password FROM users WHERE email = $1`
	user := &User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return &User{}, err
	}

	return user, nil

}
