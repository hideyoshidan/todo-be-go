package models

import "gorm.io/gorm"

// TaskHasCategory is struct for task_has_categories
type TaskHasCategory struct {
	CategoryID uint64   `gorm:"not null;index:unique_task_id_category_id,unique"`
	Category   Category `gorm:"constraint:OnDelete:CASCADE;"`
	TaskID     uint64   `gorm:"not null;index:unique_task_id_category_id,unique"`
	Task       Task     `gorm:"constraint:OnDelete:CASCADE;"`
}

// CREATE TABLE `task_has_categories` (
// 	`category_id` bigint NOT NULL,
// 	`task_id` bigint unsigned NOT NULL,
// 	UNIQUE KEY `unique_task_id_category_id` (`category_id`,`task_id`),
// 	KEY `fk_task_has_categories_task` (`task_id`),
// 	CONSTRAINT `fk_task_has_categories_category` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE CASCADE,
// 	CONSTRAINT `fk_task_has_categories_task` FOREIGN KEY (`task_id`) REFERENCES `tasks` (`id`) ON DELETE CASCADE
//   ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci

// AddCategoryIDIndexToTAskHasCategory add index to category_id
// Those struct will create above.
// Only task_id's index is created
// > https://stackoverflow.com/questions/9474954/mysql-show-create-table-shows-extra-entry-with-key
//
// Usually MySQL add index when foreign key is created. but somehow task_id has no index.
// Although MySQL creates index with foregin key, it is not added before composite index is defined
// If you want to add index, please call following in `cmd/database.go #migrate()“
//
// > https://zenn.dev/awonosuke/articles/96b3d580860e4d
// >> MySQLでは勝手にインデックスが作成されると述べたが、外部キー制約の宣言の前に、
// >> 対象カラムを含む複合キーを宣言することで、インデックスの自動生成は回避することができる模様
func AddCategoryIDIndexToTAskHasCategory(db *gorm.DB) error {
	return db.Exec("ALTER TABLE task_has_categories ADD INDEX fk_task_has_categories_category(category_id);").Error
}
