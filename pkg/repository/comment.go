package repository

import (
	"errors"

	"github.com/trungluongwww/auth/internal/model"
	"github.com/trungluongwww/auth/pkg/model/query"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Comment interface {
	Insert(doc *model.Comment) error
	Update(doc *model.Comment) error
	Delete(doc *model.Comment) error
	FirstRaw(cond *model.Comment) (*model.Comment, error)
	FindByCondition(cond query.CommentCondition, currentUserID uint32) ([]*query.CommentResult, error)
	IncrementLikeCount(commentID uint32) error
	DecrementLikeCount(commentID uint32) error
}

type comment struct {
	Tx *gorm.DB
}

func newComment(tx *gorm.DB) Comment {
	return &comment{Tx: tx}
}

func (r *comment) Insert(doc *model.Comment) error {
	return r.Tx.Create(doc).Error
}

func (r *comment) Update(doc *model.Comment) error {
	return r.Tx.Omit(clause.Associations).Select("*").Save(doc).Error
}

func (r *comment) Delete(doc *model.Comment) error {
	return r.Tx.Delete(doc).Error
}

func (r *comment) FirstRaw(cond *model.Comment) (*model.Comment, error) {
	var result *model.Comment
	err := r.Tx.Where(cond).First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}

func (r *comment) FindByCondition(cond query.CommentCondition, currentUserID uint32) ([]*query.CommentResult, error) {
	results := make([]*query.CommentResult, 0)

	tx := r.Tx.Model(&model.Comment{}).
		Preload("User")

	tx = cond.AssignID(tx)
	tx = cond.AssignName(tx)
	tx = cond.AssignPostID(tx)
	tx = cond.AssignParentID(tx)
	tx = cond.AssignOrder(tx)
	tx = cond.AssignPagination(tx)

	if err := tx.Find(&results).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return results, nil
		}
		return results, err
	}

	// Check if current user liked each comment
	if currentUserID != 0 {
		for _, result := range results {
			var count int64
			r.Tx.Model(&model.CommentLike{}).
				Where("comment_id = ? AND user_id = ?", result.ID, currentUserID).
				Count(&count)
			result.IsLiked = count > 0
		}
	}

	return results, nil
}

func (r *comment) IncrementLikeCount(commentID uint32) error {
	return r.Tx.Model(&model.Comment{}).
		Where("id = ?", commentID).
		UpdateColumn("like_count", gorm.Expr("like_count + ?", 1)).
		Error
}

func (r *comment) DecrementLikeCount(commentID uint32) error {
	return r.Tx.Model(&model.Comment{}).
		Where("id = ?", commentID).
		UpdateColumn("like_count", gorm.Expr("like_count - ?", 1)).
		Error
}
