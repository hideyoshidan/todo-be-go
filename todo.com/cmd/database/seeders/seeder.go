package seeders

import "gorm.io/gorm"

// Seeder is interface for seeders
type Seeder interface {
	seed() error
}

// Run execute seed
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
