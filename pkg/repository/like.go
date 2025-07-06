package repository

import (
	"errors"

	"github.com/trungluongwww/auth/internal/model"
	"gorm.io/gorm"
)

type PostLike interface {
	Insert(doc *model.PostLike) error
	Delete(doc *model.PostLike) error
	FirstRaw(cond *model.PostLike) (*model.PostLike, error)
	Exists(postID, userID uint32) (bool, error)
}

type CommentLike interface {
	Insert(doc *model.CommentLike) error
	Delete(doc *model.CommentLike) error
	FirstRaw(cond *model.CommentLike) (*model.CommentLike, error)
	Exists(commentID, userID uint32) (bool, error)
}

type postLike struct {
	Tx *gorm.DB
}

type commentLike struct {
	Tx *gorm.DB
}

func newPostLike(tx *gorm.DB) PostLike {
	return &postLike{Tx: tx}
}

func newCommentLike(tx *gorm.DB) CommentLike {
	return &commentLike{Tx: tx}
}

// PostLike implementations
func (r *postLike) Insert(doc *model.PostLike) error {
	return r.Tx.Create(doc).Error
}

func (r *postLike) Delete(doc *model.PostLike) error {
	return r.Tx.Delete(doc).Error
}

func (r *postLike) FirstRaw(cond *model.PostLike) (*model.PostLike, error) {
	var result *model.PostLike
	err := r.Tx.Where(cond).First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}

func (r *postLike) Exists(postID, userID uint32) (bool, error) {
	var count int64
	err := r.Tx.Model(&model.PostLike{}).
		Where("post_id = ? AND user_id = ?", postID, userID).
		Count(&count).Error
	return count > 0, err
}

// CommentLike implementations
func (r *commentLike) Insert(doc *model.CommentLike) error {
	return r.Tx.Create(doc).Error
}

func (r *commentLike) Delete(doc *model.CommentLike) error {
	return r.Tx.Delete(doc).Error
}

func (r *commentLike) FirstRaw(cond *model.CommentLike) (*model.CommentLike, error) {
	var result *model.CommentLike
	err := r.Tx.Where(cond).First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}

func (r *commentLike) Exists(commentID, userID uint32) (bool, error) {
	var count int64
	err := r.Tx.Model(&model.CommentLike{}).
		Where("comment_id = ? AND user_id = ?", commentID, userID).
		Count(&count).Error
	return count > 0, err
}
