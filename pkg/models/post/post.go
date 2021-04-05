package post

import (
	"github.com/KenLu0929/flowControlTask/pkg/models"
)

type Post struct {
	models.Basic

	UserID  uint   `gorm:"type:int" json:"user_id"`
	Title   string `gorm:"type:varchar(100)" json:"title"`
	Content string `gorm:"type:text(800)" json:"content"`
}
