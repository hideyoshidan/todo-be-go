package models

// https://stackoverflow.com/questions/9474954/mysql-show-create-table-shows-extra-entry-with-key
type TaskHasMCategory struct {
	TaskID      uint64 `gorm:"not null;index:unique_task_id_m_category_id,unique"`
	Task        Task
	MCategoryID uint32 `gorm:"not null;index:unique_task_id_m_category_id,unique"`
	MCategory   MCategory
}

// CREATE TABLE `task_has_m_categories` (
// 	`task_id` bigint unsigned NOT NULL,
// 	`m_category_id` int NOT NULL,
// 	UNIQUE KEY `unique_task_id_m_category_id` (`task_id`,`m_category_id`),
// 	KEY `fk_task_has_m_categories_m_category` (`m_category_id`),
// 	CONSTRAINT `fk_task_has_m_categories_m_category` FOREIGN KEY (`m_category_id`) REFERENCES `m_categories` (`id`),
// 	CONSTRAINT `fk_task_has_m_categories_task` FOREIGN KEY (`task_id`) REFERENCES `tasks` (`id`)
//   ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci
