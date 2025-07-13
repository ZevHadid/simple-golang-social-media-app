package repository

import (
	"simple-golang-social-media-app/internal/model"

	"gorm.io/gorm"
)

type PostRepository interface {
	FindByUserID(userID uint) ([]model.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db}
}

func (r *postRepository) FindByUserID(userID uint) ([]model.Post, error) {
	var posts []model.Post
	err := r.db.Where("user_id = ?", userID).Order("post_date desc").Find(&posts).Error
	return posts, err
}
