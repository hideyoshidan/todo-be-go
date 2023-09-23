package repositories

import (
	"gorm.io/gorm"
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
	// var mStatuses []models.MStatus
	return nil, nil
	// result := r.db.Find(&mStatuses)
	// if result.Error != nil {
	// 	log.Fatal("Error occurred while fetching data. Error : %v", result.Error)
	// 	return nil, errors.New(config.ERROR_FETCH_DATA)
	// }

	// return mStatuses, nil
}
