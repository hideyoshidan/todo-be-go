package models

import (
	"time"

	"gorm.io/gorm"
)

// User is struct for tasks users
type User struct {
	ID        uint64         `gorm:"primarykey"`
	UserName  string         `gorm:"type:varchar(255);not null"`
	Email     string         `gorm:"type:varchar(255);not null;index:unique_email_deleted_at,unique"`
	Passowrd  string         `gorm:"type:varchar(255);not null;"`
	CreatedAt time.Time      `gorm:"type:datetime(6);not null"`
	UpdatedAt time.Time      `gorm:"type:datetime(6)"`
	DeletedAt gorm.DeletedAt `gorm:"index:unique_email_deleted_at,unique"`
}

// CREATE TABLE `users` (
// 	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
// 	`user_name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
// 	`email` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
// 	`passowrd` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
// 	`created_at` datetime(6) NOT NULL,
// 	`updated_at` datetime(6) DEFAULT NULL,
// 	`deleted_at` datetime(6) DEFAULT NULL,
// 	PRIMARY KEY (`id`),
// 	UNIQUE KEY `unique_email_deleted_at` (`email`,`deleted_at`)
//   ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci
