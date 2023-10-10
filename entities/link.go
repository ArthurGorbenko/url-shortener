package entities

import (
	"time"
)

type ShortLink struct {
	Origin    string    `gorm:"unique;type:varchar(255)"`
	Short     string    `gorm:"unique;type:varchar(255)"`
	CreatedAt time.Time `gorm:"type:timestamp"`
}
