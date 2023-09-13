package db

import (
	"benjaminlee.dev/ProjectManagement/models" // Import your models

	"gorm.io/gorm"
)

func CreateBoard(db *gorm.DB, board *models.Board) error {
	return db.Create(board).Error
}

func GetBoardByID(db *gorm.DB, boardID uint) (*models.Board, error) {
	var board models.Board
	if err := db.First(&board, boardID).Error; err != nil {
		return nil, err
	}
	return &board, nil
}

func UpdateBoard(db *gorm.DB, board *models.Board) error {
	return db.Save(board).Error
}

func DeleteBoard(db *gorm.DB, boardID uint) error {
	return db.Delete(&models.Board{}, boardID).Error
}
