package database

import (
	"log"

	"gorm.io/gorm"
	"todo.com/cmd/database/seeders"
	"todo.com/config"
	"todo.com/domain/models"
	"todo.com/infrastructure/database"
)

const (
	DB_TYPE_MIGRATE = "migrate"
	DB_TYPE_DROP    = "drop"
	DB_SEED         = "seed"
)

// Gorm is struct for *gorm.DB
type Gorm struct {
	*gorm.DB
}

// Run execute DB operation
func Run(operate_type string) error {
	db, err := newDB()
	if err != nil {
		log.Fatalf("Failed to establish DB connection : %v", err)
		return err
	}

	err = nil
	switch operate_type {
	case DB_TYPE_MIGRATE:
		err = db.migrate()
	case DB_TYPE_DROP:
		err = db.dropTable()
	case DB_SEED:
		err = seeders.Run(db.DB)
	}

	if err != nil {
		return err
	}

	return nil
}

// neDB generate instance of Gorm
func newDB() (*Gorm, error) {
	dbConn, err := database.NewDBConn(
		database.NewWithNoOpt(
			config.DB_USER,
			config.DB_PASSWORD,
			config.DB_ADDR,
			config.DB_NAME,
			config.LOCATION,
		),
	)

	if err != nil {
		return nil, err
	}
	return &Gorm{dbConn}, nil
}

// migrate is to create Tables
func (g *Gorm) migrate() error {
	return g.AutoMigrate(
		&models.Category{},
		&models.MStatus{},
		&models.User{},
		&models.Task{},
		&models.TaskHasCategory{},
	)
}

// dropTable is to delete tables
func (g *Gorm) dropTable() error {
	return g.Migrator().DropTable(
		&models.Category{},
		&models.MStatus{},
		&models.User{},
		&models.Task{},
		&models.TaskHasCategory{},
	)
}
