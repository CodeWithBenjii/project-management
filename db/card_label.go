package db

import (
	"benjaminlee.dev/ProjectManagement/models" // Import your models

	"gorm.io/gorm"
)

func CreateCardLabel(db *gorm.DB, cardLabel *models.CardLabel) error {
	return db.Create(cardLabel).Error
}

func GetCardLabelByID(db *gorm.DB, cardLabelID uint) (*models.CardLabel, error) {
	var cardLabel models.CardLabel
	if err := db.First(&cardLabel, cardLabelID).Error; err != nil {
		return nil, err
	}
	return &cardLabel, nil
}

func UpdateCardLabel(db *gorm.DB, cardLabel *models.CardLabel) error {
	return db.Save(cardLabel).Error
}

func DeleteCardLabel(db *gorm.DB, cardLabelID uint) error {
	return db.Delete(&models.CardLabel{}, cardLabelID).Error
}
