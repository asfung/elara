package repositories

import (
	"errors"

	"github.com/asfung/elara/database"
	"github.com/asfung/elara/internal/entities"
	"gorm.io/gorm"
)

type userPostgresRepository struct {
	*BaseRepository[entities.User]
}

func NewUserPostgresRepository(db database.Database) UserRepository {
	return &userPostgresRepository{
		BaseRepository: NewBaseRepository[entities.User](db),
	}
}

func (r *userPostgresRepository) FindByEmail(email string) (entities.User, error) {
	var user entities.User
	if err := r.db.GetDb().Where("email = ?", email).First(&user).Error; err != nil {
		// if errors.Is(err, gorm.ErrRecordNotFound) {
		// 	return entities.User{}, errors.New("user not found")
		// }
		return entities.User{}, err
	}
	return user, nil
}

func (r *userPostgresRepository) FindByRefreshToken(refreshToken string) (entities.User, error) {
	var user entities.User
	if err := r.db.GetDb().Where("refresh_token = ?", refreshToken).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.User{}, errors.New("user not found")
		}
		return entities.User{}, err
	}
	return user, nil
}

func (r *userPostgresRepository) FindByUserId(userId string) (entities.User, error) {
	var user entities.User
	if err := r.db.GetDb().Where("user_id = ?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.User{}, errors.New("user not found")
		}
		return entities.User{}, err
	}
	return user, nil
}

func (r *userPostgresRepository) FindByProvider(provider string, providerUserId string) (entities.User, error) {
	var user entities.User
	if err := r.db.GetDb().Where("provider = ? AND provider_user_id = ?", provider, providerUserId).First(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}
