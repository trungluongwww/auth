package query

import (
	"fmt"

	"github.com/trungluongwww/auth/internal/model"
	"gorm.io/gorm"
)

type PostResult struct {
	model.Post
	User    *model.User `json:"user" gorm:"foreignKey:UserID;references:ID"`
	IsLiked bool        `json:"is_liked"`
}

type CommentResult struct {
	model.Comment
	User    *model.User `json:"user" gorm:"foreignKey:UserID;references:ID"`
	IsLiked bool        `json:"is_liked"`
}

type PostCondition struct {
	CommonCondition
	UserID   uint32
	IsPublic *bool
	Search   string
}

func (c *PostCondition) AssignUserID(tx *gorm.DB) *gorm.DB {
	if c.UserID != 0 {
		tx = tx.Where("user_id = ?", c.UserID)
	}
	return tx
}

func (c *PostCondition) AssignIsPublic(tx *gorm.DB) *gorm.DB {
	if c.IsPublic != nil {
		tx = tx.Where("is_public = ?", *c.IsPublic)
	}
	return tx
}

func (c *PostCondition) AssignSearch(tx *gorm.DB) *gorm.DB {
	if c.Search != "" {
		tx = tx.Where("(title LIKE ? OR content LIKE ?)",
			fmt.Sprintf("%%%s%%", c.Search),
			fmt.Sprintf("%%%s%%", c.Search))
	}
	return tx
}

type CommentCondition struct {
	CommonCondition
	PostID   uint32
	ParentID *uint32
}

func (c *CommentCondition) AssignPostID(tx *gorm.DB) *gorm.DB {
	if c.PostID != 0 {
		tx = tx.Where("post_id = ?", c.PostID)
	}
	return tx
}

func (c *CommentCondition) AssignParentID(tx *gorm.DB) *gorm.DB {
	if c.ParentID != nil {
		tx = tx.Where("parent_id = ?", *c.ParentID)
	} else {
		tx = tx.Where("parent_id IS NULL")
	}
	return tx
}
