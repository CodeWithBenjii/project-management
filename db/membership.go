package db

import (
	"benjaminlee.dev/ProjectManagement/models" // Import your models

	"gorm.io/gorm"
)

func CreateMembership(db *gorm.DB, membership *models.Membership) error {
	return db.Create(membership).Error
}

func GetMembershipByID(db *gorm.DB, membershipID uint) (*models.Membership, error) {
	var membership models.Membership
	if err := db.First(&membership, membershipID).Error; err != nil {
		return nil, err
	}
	return &membership, nil
}

func UpdateMembership(db *gorm.DB, membership *models.Membership) error {
	return db.Save(membership).Error
}

func DeleteMembership(db *gorm.DB, membershipID uint) error {
	return db.Delete(&models.Membership{}, membershipID).Error
}
