package repository

import (
	"errors"

	"github.com/trungluongwww/auth/internal/model"
	"github.com/trungluongwww/auth/pkg/model/query"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Post interface {
	Insert(doc *model.Post) error
	Update(doc *model.Post) error
	Delete(doc *model.Post) error
	FirstRaw(cond *model.Post) (*model.Post, error)
	FindByCondition(cond query.PostCondition, currentUserID uint32) ([]*query.PostResult, error)
	IncrementViewCount(postID uint32) error
	IncrementLikeCount(postID uint32) error
	DecrementLikeCount(postID uint32) error
	IncrementCommentCount(postID uint32) error
	DecrementCommentCount(postID uint32) error
}

type post struct {
	Tx *gorm.DB
}

func newPost(tx *gorm.DB) Post {
	return &post{Tx: tx}
}

func (r *post) Insert(doc *model.Post) error {
	return r.Tx.Create(doc).Error
}

func (r *post) Update(doc *model.Post) error {
	return r.Tx.Omit(clause.Associations).Select("*").Save(doc).Error
}

func (r *post) Delete(doc *model.Post) error {
	return r.Tx.Delete(doc).Error
}

func (r *post) FirstRaw(cond *model.Post) (*model.Post, error) {
	var result *model.Post
	err := r.Tx.Where(cond).First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}

func (r *post) FindByCondition(cond query.PostCondition, currentUserID uint32) ([]*query.PostResult, error) {
	results := make([]*query.PostResult, 0)

	tx := r.Tx.Model(&model.Post{}).
		Preload("User")

	tx = cond.AssignID(tx)
	tx = cond.AssignName(tx)
	tx = cond.AssignUserID(tx)
	tx = cond.AssignIsPublic(tx)
	tx = cond.AssignSearch(tx)
	tx = cond.AssignOrder(tx)
	tx = cond.AssignPagination(tx)

	if err := tx.Find(&results).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return results, nil
		}
		return results, err
	}

	// Check if current user liked each post
	if currentUserID != 0 {
		for _, result := range results {
			var count int64
			r.Tx.Model(&model.PostLike{}).
				Where("post_id = ? AND user_id = ?", result.ID, currentUserID).
				Count(&count)
			result.IsLiked = count > 0
		}
	}

	return results, nil
}

func (r *post) IncrementViewCount(postID uint32) error {
	return r.Tx.Model(&model.Post{}).
		Where("id = ?", postID).
		UpdateColumn("view_count", gorm.Expr("view_count + ?", 1)).
		Error
}

func (r *post) IncrementLikeCount(postID uint32) error {
	return r.Tx.Model(&model.Post{}).
		Where("id = ?", postID).
		UpdateColumn("like_count", gorm.Expr("like_count + ?", 1)).
		Error
}

func (r *post) DecrementLikeCount(postID uint32) error {
	return r.Tx.Model(&model.Post{}).
		Where("id = ?", postID).
		UpdateColumn("like_count", gorm.Expr("like_count - ?", 1)).
		Error
}

func (r *post) IncrementCommentCount(postID uint32) error {
	return r.Tx.Model(&model.Post{}).
		Where("id = ?", postID).
		UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).
		Error
}

func (r *post) DecrementCommentCount(postID uint32) error {
	return r.Tx.Model(&model.Post{}).
		Where("id = ?", postID).
		UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).
		Error
}
