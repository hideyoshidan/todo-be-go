package seeders

import (
	"log"
	"time"

	"gorm.io/gorm"
	"todo.com/domain/models"
)

// MStatusSeeder is struct for m_statues
type MStatusSeeder struct {
	*gorm.DB
}

// NewMStatusSeeder return struct of MStatusSeeder
func NewMStatusSeeder(gorm *gorm.DB) *MStatusSeeder {
	return &MStatusSeeder{
		gorm,
	}
}

func (s *MStatusSeeder) seed() error {
	statuses := []models.MStatus{
		{
			Name:      "TODO",
			ColorCode: "FB0000",
			CreatedAt: time.Now(),
		},
		{
			Name:      "DOING",
			ColorCode: "00B0FB",
			CreatedAt: time.Now(),
		},
		{
			Name:      "DONE",
			ColorCode: "00C738",
			CreatedAt: time.Now(),
		},
	}
	r := s.Create(statuses)
	if r.Error != nil {
		return r.Error
	}

	log.Printf("Seed successed. Inserted item count is : %v", r.RowsAffected)
	return nil
}

// seed is implemented Seeder interface
// func (s *MStatusSeeder) seed() error {
// 	statuses := []models.MStatus{
// 		{
// 			Name:      "TODO",
// 			ColorCode: "FB0000",
// 		},
// 		{
// 			Name:      "DOING",
// 			ColorCode: "00B0FB",
// 		},
// 		{
// 			Name:      "DONE",
// 			ColorCode: "00C738",
// 		},
// 	}

// 	sql := "INSERT INTO m_statuses (`name`, `color_code`, `created_at`) VALUES %s"
// 	var values []string
// 	for i, status := range statuses {
// 		var icon string
// 		if i == len(statuses)-1 {
// 			icon = ";"
// 		} else {
// 			icon = ","
// 		}
// 		values = append(values, fmt.Sprintf("('%s', '%s', '%s')%s", status.Name, status.ColorCode, time.Now().Format(time.RFC3339Nano), icon))
// 	}

// 	sql = fmt.Sprintf(sql, strings.Join(values, ""))
// 	return s.DB.Exec(sql).Error
// }
