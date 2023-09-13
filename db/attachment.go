package db

import (
	"benjaminlee.dev/ProjectManagement/models" // Import your models
	"gorm.io/gorm"
)

func CreateAttachment(db *gorm.DB, attachment *models.Attachment) error {
	return db.Create(attachment).Error
}

func GetAttachmentByID(db *gorm.DB, attachmentID uint) (*models.Attachment, error) {
	var attachment models.Attachment
	if err := db.First(&attachment, attachmentID).Error; err != nil {
		return nil, err
	}
	return &attachment, nil
}

func UpdateAttachment(db *gorm.DB, attachment *models.Attachment) error {
	return db.Save(attachment).Error
}

func DeleteAttachment(db *gorm.DB, attachmentID uint) error {
	return db.Delete(&models.Attachment{}, attachmentID).Error
}
