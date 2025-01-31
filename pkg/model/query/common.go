package query

import (
	"fmt"
	"gorm.io/gorm"
)

const (
	CreatedAtColumn = "created_at"
	SortDescending  = "desc"
	SortAscending   = "asc"
)

type Pagination struct {
	Limit int
	Page  int
}

type Order struct {
	OrderBy    string
	OrderValue string
}

type CommonCondition struct {
	ID         int
	Name       string
	Order      *Order
	Pagination *Pagination
}

func (c *CommonCondition) AssignID(tx *gorm.DB) *gorm.DB {
	if c.ID != 0 {
		tx = tx.Where("id = ?", c.ID)
	}

	return tx
}

func (c *CommonCondition) AssignName(tx *gorm.DB) *gorm.DB {
	if c.Name != "" {
		return tx.Where("name like ?", fmt.Sprintf("%%%s%%%", c.Name))
	}

	return tx
}

func (c *CommonCondition) AssignOrder(tx *gorm.DB) *gorm.DB {
	if c.Order != nil {
		if c.Order.OrderBy == "" {
			c.Order.OrderBy = CreatedAtColumn
		}

		if c.Order.OrderValue == "" {
			c.Order.OrderValue = SortDescending
		}
		return tx.Order(fmt.Sprintf("%s %s", c.Order.OrderBy, c.Order.OrderValue))
	}
	return tx
}

func (c *CommonCondition) AssignPagination(tx *gorm.DB) *gorm.DB {
	if c.Pagination != nil {
		if c.Pagination.Limit == 0 {
			c.Pagination.Limit = 50
		}

		if c.Pagination.Page < 1 {
			c.Pagination.Page = 1
		}

		return tx.Limit(c.Pagination.Limit).Offset((c.Pagination.Page - 1) * c.Pagination.Limit)
	}

	return tx
}
