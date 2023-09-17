package database

import (
	"log"

	"gorm.io/gorm"
	"todo.com/config"
	"todo.com/domain/models"
	"todo.com/infrastructure/database"
)

type Gorm struct {
	*gorm.DB
}

func Run(operate_type string) error {
	db, err := newDB()
	if err != nil {
		log.Fatalf("Failed to establish DB connection : %v", err)
		return err
	}

	err = nil
	switch operate_type {
	case "migrate":
		err = db.migrate()
	case "drop":
		err = db.dropTable()
	}

	if err != nil {
		return err
	}

	return nil
}

func newDB() (*Gorm, error) {
	dbConn, err := database.NewDBConn(
		database.NewWithOpt(
			config.DB_USER,
			config.DB_PASSWORD,
			config.DB_ADDR,
			config.DB_NAME,
			config.LOCATION,
			&gorm.Config{
				DisableForeignKeyConstraintWhenMigrating: true,
			},
		),
	)

	if err != nil {
		return nil, err
	}
	return &Gorm{dbConn}, nil
}

func (g *Gorm) migrate() error {
	return g.AutoMigrate(&models.MCategories{})
}

func (g *Gorm) dropTable() error {
	return g.Migrator().DropTable(&models.MCategories{})
}
