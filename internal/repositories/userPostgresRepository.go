package repositories

import (
	"errors"

	"github.com/asfung/elara/database"
	"github.com/asfung/elara/internal/entities"
	"gorm.io/gorm"
)

type userPostgresRepository struct {
	db database.Database
}

func NewUserPostgresRepository(db database.Database) UserRepository {
	return &userPostgresRepository{
		db: db,
	}
}

func (r *userPostgresRepository) Create(user entities.User) (entities.User, error) {
	if err := r.db.GetDb().Create(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (r *userPostgresRepository) Update(user entities.User) (entities.User, error) {
	if err := r.db.GetDb().Save(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (r *userPostgresRepository) FindById(id uint32) (entities.User, error) {
	var user entities.User
	if err := r.db.GetDb().First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.User{}, errors.New("user not found")
		}
		return entities.User{}, err
	}
	return user, nil
}

func (r *userPostgresRepository) FindByEmail(email string) (entities.User, error) {
	var user entities.User
	if err := r.db.GetDb().Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.User{}, errors.New("user not found")
		}
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

func (r *userPostgresRepository) Delete(id uint32) error {
	if err := r.db.GetDb().Delete(&entities.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
