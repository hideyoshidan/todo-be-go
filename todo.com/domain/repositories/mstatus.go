package repositories

import (
	"errors"

	"gorm.io/gorm"
	"todo.com/config"
	"todo.com/domain/models"
)

type MStatusRespository struct {
	db *gorm.DB
}

func NewMStatusRespository(db *gorm.DB) *MStatusRespository {
	return &MStatusRespository{
		db: db,
	}
}

// FetchAll fetche all status data
func (r *MStatusRespository) FetchAll() ([]models.MStatus, error) {
	var mStatuses []models.MStatus
	result := r.db.Find(&mStatuses)
	if result.Error != nil {
		return nil, errors.New(config.ERROR_FETCH_DATA)
	}

	return mStatuses, nil
}
