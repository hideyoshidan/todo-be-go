package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          uint64 `gorm:"primarykey"`
	UserID      uint64
	User        User
	Title       string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text;"`
	Deadline    time.Time `gorm:"type:date;"`
	MStatusID   uint32
	MStatus     MStatus
	CreatedAt   time.Time `gorm:"type:datetime(6);not null"`
	UpdatedAt   time.Time `gorm:"type:datetime(6)"`
	DeletedAt   gorm.DeletedAt
}

// CREATE TABLE `tasks` (
// 	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
// 	`user_id` bigint unsigned DEFAULT NULL,
// 	`title` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
// 	`description` text COLLATE utf8mb4_general_ci,
// 	`deadline` date DEFAULT NULL,
// 	`m_status_id` int DEFAULT NULL,
// 	`created_at` datetime(6) NOT NULL,
// 	`updated_at` datetime(6) DEFAULT NULL,
// 	`deleted_at` datetime(6) DEFAULT NULL,
// 	PRIMARY KEY (`id`),
// 	KEY `fk_tasks_user` (`user_id`),
// 	KEY `fk_tasks_m_status` (`m_status_id`),
// 	CONSTRAINT `fk_tasks_m_status` FOREIGN KEY (`m_status_id`) REFERENCES `m_statuses` (`id`),
// 	CONSTRAINT `fk_tasks_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
//   ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci
