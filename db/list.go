package db

import (
	"benjaminlee.dev/ProjectManagement/models" // Import your models

	"gorm.io/gorm"
)

func CreateList(db *gorm.DB, list *models.List) error {
	return db.Create(list).Error
}

func GetListByID(db *gorm.DB, listID uint) (*models.List, error) {
	var list models.List
	if err := db.First(&list, listID).Error; err != nil {
		return nil, err
	}
	return &list, nil
}

func UpdateList(db *gorm.DB, list *models.List) error {
	return db.Save(list).Error
}

func DeleteList(db *gorm.DB, listID uint) error {
	return db.Delete(&models.List{}, listID).Error
}
