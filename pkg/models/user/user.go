package user

import (
	"time"

	"github.com/KenLu0929/flowControlTask/pkg/models"
)

type User struct {
	models.Basic

	Account  string     `gorm:"type:varchar(40)" json:"account"`
	Password string     `gorm:"type:varchar(100)" json:"password"`
	Username string     `gorm:"type:varchar(20)" json:"username"`
	Gender   string     `gorm:"type:varchar(5);NULL;DEFAULT:NULL" json:"gender"`
	Birthday *time.Time `gorm:"type:date;NULL;DEFAULT:NULL" json:"birthday"`
}
