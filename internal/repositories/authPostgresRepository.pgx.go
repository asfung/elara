package repositories

import (
	"context"
	"errors"

	"github.com/asfung/elara/database"
	"github.com/asfung/elara/internal/entities"
	"github.com/jackc/pgx/v5"
)

type authPgxRepository struct {
	db database.Database
}

func NewAuthPgxRepository(db database.Database) AuthRepository {
	return &authPgxRepository{db: db}
}

func (r *authPgxRepository) FindAccessToken(accessToken string) (entities.User, error) {
	var user entities.User
	query := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE access_token = $1"
	err := r.db.GetDb().QueryRow(context.Background(), query, accessToken).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return entities.User{}, errors.New("user not found")
		}
		return entities.User{}, err
	}
	return user, nil
}
