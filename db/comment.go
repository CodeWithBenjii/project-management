package db

import (
	"benjaminlee.dev/ProjectManagement/models" // Import your models

	"gorm.io/gorm"
)

func CreateComment(db *gorm.DB, comment *models.Comment) error {
	return db.Create(comment).Error
}

func GetCommentByID(db *gorm.DB, commentID uint) (*models.Comment, error) {
	var comment models.Comment
	if err := db.First(&comment, commentID).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func UpdateComment(db *gorm.DB, comment *models.Comment) error {
	return db.Save(comment).Error
}

func DeleteComment(db *gorm.DB, commentID uint) error {
	return db.Delete(&models.Comment{}, commentID).Error
}
