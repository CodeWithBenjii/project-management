package db

import (
	"benjaminlee.dev/ProjectManagement/models" // Import your models

	"gorm.io/gorm"
)

func CreateCard(db *gorm.DB, card *models.Card) error {
	return db.Create(card).Error
}

func GetCardByID(db *gorm.DB, cardID uint) (*models.Card, error) {
	var card models.Card
	if err := db.First(&card, cardID).Error; err != nil {
		return nil, err
	}
	return &card, nil
}

func UpdateCard(db *gorm.DB, card *models.Card) error {
	return db.Save(card).Error
}

func DeleteCard(db *gorm.DB, cardID uint) error {
	return db.Delete(&models.Card{}, cardID).Error
}
