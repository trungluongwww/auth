package db

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	Seeder interface {
		Run() error
	}
	initialSeed struct {
		gdb *gorm.DB
	}
)

func NewInitialSeeder(gdb *gorm.DB) Seeder {
	return &initialSeed{gdb: gdb}
}

func (s *initialSeed) Run() error {
	seeds := append(
		make([]interface{}, 0),
	)
	return s.gdb.Transaction(func(tx *gorm.DB) error {
		for _, seed := range seeds {
			err := s.gdb.
				Clauses(clause.OnConflict{DoNothing: true}).
				Create(seed).
				Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}
