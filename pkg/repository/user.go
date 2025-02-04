package repository

import (
	"errors"
	"github.com/trungluongwww/auth/internal/model"
	"github.com/trungluongwww/auth/pkg/model/query"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User interface {
	Insert(doc *model.User) error
	Update(doc *model.User) error
	Delete(doc *model.User) error
	FirstRaw(cond *model.User) (*model.User, error)
	FindByCondition(cond query.CommonCondition) ([]*query.UserResult, error)

	InsertUserFacebookLogin(doc *model.UserFacebookLogin) error
}

type user struct {
	Tx *gorm.DB
}

func newUser(tx *gorm.DB) User {
	return &user{Tx: tx}
}

func (r *user) Insert(doc *model.User) error {
	return r.Tx.Create(doc).Error
}

func (r *user) Update(doc *model.User) error {
	return r.Tx.Omit(clause.Associations).Select("*").Save(doc).Error
}

func (r *user) Delete(doc *model.User) error {
	return r.Tx.Delete(doc).Error
}

func (r *user) FirstRaw(cond *model.User) (*model.User, error) {
	var result *model.User
	err := r.Tx.Where(cond).First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

func (r *user) FindByCondition(cond query.CommonCondition) ([]*query.UserResult, error) {
	results := make([]*query.UserResult, 0)

	tx := r.Tx.Model(&model.User{})

	tx = cond.AssignID(tx)
	tx = cond.AssignName(tx)
	tx = cond.AssignOrder(tx)
	tx = cond.AssignPagination(tx)

	if err := tx.Find(&results).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return results, nil
		}
		return results, err
	}

	return results, nil
}

func (r *user) InsertUserFacebookLogin(doc *model.UserFacebookLogin) error {
	return r.Tx.Create(doc).Error
}
