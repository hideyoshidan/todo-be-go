package models

import (
	"time"
)

// m_statuses tables
type MStatus struct {
	ID        uint32    `gorm:"type:int;primarykey"`
	Name      string    `gorm:"type:varchar(255);not null"`
	ColorCode string    `gorm:"type:varchar(255);unique;not null"`
	CreatedAt time.Time `gorm:"type:datetime(6);not null"`
}

// CREATE TABLE `m_statuses` (
// 	`id` int NOT NULL AUTO_INCREMENT,
// 	`name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
// 	`color_code` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
// 	`created_at` datetime(6) NOT NULL,
// 	PRIMARY KEY (`id`),
// 	UNIQUE KEY `color_code` (`color_code`)
//   ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci
