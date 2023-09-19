package seeders

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
	"todo.com/domain/models"
)

type MStatusSeeder struct {
	*gorm.DB
}

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

	sql := "INSERT INTO m_status (`name`, `color_code`, `created_at`) VALUES %s"
	var values []string
	for i, status := range statuses {
		var icon string
		if i == len(statuses)-1 {
			icon = ";"
		} else {
			icon = ","
		}
		values = append(values, fmt.Sprintf("('%s', '%s', '%s')%s", status.Name, status.ColorCode, status.CreatedAt, icon))
	}

	sql = fmt.Sprintf(sql, strings.Join(values, ""))
	fmt.Printf("%v", sql)
	// ここから！！！！！！！！
	return nil
}
