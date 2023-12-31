package models

import (
	"time"
)

// Category is struct for categories
type Category struct {
	ID        uint64 `gorm:"type:int;primarykey"`
	UserID    uint64
	User      User
	Name      string    `gorm:"type:varchar(255);not null"`
	ColorCode string    `gorm:"type:varchar(255);unique;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime"`
}

// CREATE TABLE `categories` (
// 	`id` bigint NOT NULL AUTO_INCREMENT,
// 	`user_id` bigint unsigned DEFAULT NULL,
// 	`name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
// 	`color_code` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
// 	`created_at` datetime NOT NULL,
// 	`updated_at` datetime DEFAULT NULL,
// 	PRIMARY KEY (`id`),
// 	UNIQUE KEY `color_code` (`color_code`),
// 	KEY `fk_categories_user` (`user_id`),
// 	CONSTRAINT `fk_categories_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci
