package seeders

import "gorm.io/gorm"

type Seeder interface {
	seed() error
}

func Run(gorm *gorm.DB) error {
	seeders := []Seeder{
		NewMStatusSeeder(gorm),
	}

	for _, model := range seeders {
		if err := seeds(model); err != nil {
			return err
		}
	}

	return nil
}

func seeds(s Seeder) error {
	return s.seed()
}
