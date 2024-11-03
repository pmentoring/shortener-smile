package repository

import (
	"database/sql"
	"shortener-smile/internal/auth/model"
)

type UserRepository interface {
	GetUserByLogin(login string) (*model.User, error)
}

type PostgresUserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (r PostgresUserRepository) GetUserByLogin(login string) (*model.User, error) {
	row := r.db.QueryRow(`SELECT * FROM public.user WHERE login = $1`, login)

	err := row.Err()
	if err != nil {
		return nil, err
	}

	user := new(model.User)

	err = row.Scan(&user.Id, &user.Login, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}
