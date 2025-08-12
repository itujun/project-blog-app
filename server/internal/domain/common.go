// Helper Timestamp
package domain

import (
	"time"

	"gorm.io/gorm"
)

// TimeFields mengelompokkan field waktu standar.
type TimeFields struct {
	CreatedAt time.Time 	`json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time 	`json:"updated_at,omitempty" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}