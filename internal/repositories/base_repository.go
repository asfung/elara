package repositories

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/asfung/elara/database"
	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	db database.Database
}

func NewBaseRepository[T any](db database.Database) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

func (r *BaseRepository[T]) Create(entity T) (T, error) {
	if err := r.db.GetDb().Create(&entity).Error; err != nil {
		return entity, err
	}
	return entity, nil
}

func (r *BaseRepository[T]) Update(entity T) (T, error) {
	if err := r.db.GetDb().Save(&entity).Error; err != nil {
		return entity, err
	}
	return entity, nil
}

func (r *BaseRepository[T]) FindById(id any) (*T, error) {
	var entity T
	// if err := r.db.GetDb().First(&entity, id).Error; err != nil {
	if err := r.db.GetDb().First(&entity, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			typeName := reflect.TypeOf(entity).Name()
			if typeName == "" {
				typeName = "record"
			}
			return nil, fmt.Errorf("%s not found", typeName)
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
