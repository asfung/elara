package repositories

import (
	"context"
	"errors"

	"github.com/asfung/elara/database"
	"github.com/asfung/elara/internal/entities"
	"github.com/jackc/pgx/v5"
)

type userPgxRepository struct {
	db database.Database
}

func NewUserPgxRepository(db database.Database) UserRepository {
	return &userPgxRepository{db: db}
}

func (r *userPgxRepository) FindByEmail(email string) (entities.User, error) {
	var user entities.User
	query := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = $1"
	err := r.db.GetDb().QueryRow(context.Background(), query, email).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return entities.User{}, errors.New("user not found")
		}
		return entities.User{}, err
	}
	return user, nil
}

func (r *userPgxRepository) FindById(id string) (entities.User, error) {
	panic("implement me")
}

func (r *userPgxRepository) FindByRefreshToken(refreshToken string) (entities.User, error) {
	panic("implement me")
}

func (r *userPgxRepository) Create(user entities.User) (entities.User, error) {
	panic("implement me")
}

func (r *userPgxRepository) Update(user entities.User) (entities.User, error) {
	panic("implement me")
}
