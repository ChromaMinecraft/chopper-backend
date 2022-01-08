package domain

import (
	"time"

	"gorm.io/gorm"
)

type Report struct {
	ID          int            `json:"id,omitempty" gorm:"column:id"`
	CreatedAt   time.Time      `json:"-" gorm:"column:created_at"`
	UpdatedAt   time.Time      `json:"-" gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
	ReporterID  int64          `json:"reporter_id" gorm:"column:reporter_id"`
	ReportedID  int64          `json:"reported_id" gorm:"column:reported_id"`
	ReportID    string         `json:"report_id" gorm:"column:report_id"`
	Description string         `json:"description" gorm:"column:description"`
}

func (Report) TableName() string {
	return "report"
}
