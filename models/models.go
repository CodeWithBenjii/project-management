package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username      string `gorm:"unique"`
    Email         string `gorm:"unique"`
    PasswordHash  string
    ProfilePictureURL string
}

type Board struct {
    gorm.Model
    Title       string
    Description string
    CreatorUserID uint
    CreatorUser   User `gorm:"foreignKey:CreatorUserID"`
}

type List struct {
    gorm.Model
    Title    string
    Position int
    BoardID  uint
    Board    Board `gorm:"foreignKey:BoardID"`
}

type Card struct {
    gorm.Model
    Title       string
    Description string
    Position    int
    ListID      uint
    List        List `gorm:"foreignKey:ListID"`
}

type Label struct {
    gorm.Model
    Name  string
    Color string
}

type CardLabel struct {
    gorm.Model
    CardID  uint
    LabelID uint
}

type Comment struct {
    gorm.Model
    Text     string
    CardID   uint
    UserID   uint
    User     User `gorm:"foreignKey:UserID"`
    Timestamp time.Time
}

type Attachment struct {
    gorm.Model
    FileName   string
    URL        string
    CardID     uint
    UserID     uint
    User       User `gorm:"foreignKey:UserID"`
    Timestamp  time.Time
}

type Membership struct {
    gorm.Model
    BoardID uint
    UserID  uint
}
