package db

import (
	"benjaminlee.dev/ProjectManagement/models" // Import your models

	"gorm.io/gorm"
)

func CreateLabel(db *gorm.DB, label *models.Label) error {
	return db.Create(label).Error
}

func GetLabelByID(db *gorm.DB, labelID uint) (*models.Label, error) {
	var label models.Label
	if err := db.First(&label, labelID).Error; err != nil {
		return nil, err
	}
	return &label, nil
}

func UpdateLabel(db *gorm.DB, label *models.Label) error {
	return db.Save(label).Error
}

func DeleteLabel(db *gorm.DB, labelID uint) error {
	return db.Delete(&models.Label{}, labelID).Error
}
