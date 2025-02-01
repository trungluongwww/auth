package repository

import (
	"gorm.io/gorm"
)

type Repository interface {
	NewTransaction(fc func(tx Repository) error) error
	NewUser() User
	NewAccount() Account
}

type repositoryIml struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryIml{db: db}
}

func (r *repositoryIml) NewTransaction(fc func(tx Repository) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		return fc(NewRepository(tx))
	})
}

func (r *repositoryIml) NewUser() User {
	return newUser(r.db)
}

func (r *repositoryIml) NewAccount() Account {
	return newAccount(r.db)
}
