package repositories

import (
	"errors"

	"github.com/asfung/elara/database"
	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	db database.Database
}

func NewBaseRepository[T any](db database.Database) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

func (r *BaseRepository[T]) Create(entity *T) error {
	return r.db.GetDb().Create(entity).Error
}

func (r *BaseRepository[T]) Update(entity *T) error {
	return r.db.GetDb().Save(entity).Error
}

func (r *BaseRepository[T]) FindById(id any) (*T, error) {
	var entity T
	if err := r.db.GetDb().First(&entity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("record not found")
		}
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepository[T]) Delete(id any) error {
	var entity T
	if err := r.db.GetDb().Delete(&entity, id).Error; err != nil {
		return err
	}
	return nil
}
