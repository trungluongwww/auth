package repository

import (
	"errors"
	"fmt"
	"github.com/trungluongwww/auth/internal/model"
	"github.com/trungluongwww/auth/pkg/model/query"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Account interface {
	Insert(doc *model.Account) error
	Update(doc *model.Account) error
	FirstRaw(cond *model.Account) (*model.Account, error)
	FirstByRefreshToken(refreshToken string) (*query.AccountResult, error)

	// account refresh token
	InsertAccountRefreshToken(doc *model.AccountRefreshToken) error
	UpdateAccountRefreshToken(doc *model.AccountRefreshToken) error
}

type account struct {
	Tx *gorm.DB
}

func newAccount(tx *gorm.DB) Account {
	return &account{Tx: tx}
}

func (r *account) FirstRaw(cond *model.Account) (*model.Account, error) {
	var result *model.Account
	err := r.Tx.Where(cond).First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

func (r *account) Insert(doc *model.Account) error {
	return r.Tx.Omit(clause.Associations).Create(doc).Error
}

func (r *account) Update(doc *model.Account) error {
	return r.Tx.Omit(clause.Associations).Updates(doc).Error
}

func (r *account) InsertAccountRefreshToken(doc *model.AccountRefreshToken) error {
	return r.Tx.Omit(clause.Associations).Create(doc).Error
}

func (r *account) UpdateAccountRefreshToken(doc *model.AccountRefreshToken) error {
	return r.Tx.Omit(clause.Associations).Updates(doc).Error
}

func (r *account) FirstByRefreshToken(refreshToken string) (*query.AccountResult, error) {
	tx := r.Tx.Table(model.TableNameAccount).
		Joins(fmt.Sprintf("INNER JOIN %s as art on art.account_id = account.id", model.TableNameAccountRefreshToken)).
		Where("art.token = ?", refreshToken).
		Preload("User")

	result := &query.AccountResult{}
	if err := tx.First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}
