package mysql

import (
	"context"

	"github.com/ChromaMinecraft/chopper-backend/internal/core/domain"
	"github.com/ChromaMinecraft/chopper-backend/internal/core/ports"
	"gorm.io/gorm"
)

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ports.ReportRepositorier {
	return &reportRepository{db: db}
}

func (r *reportRepository) GetAll(ctx context.Context) ([]domain.Report, error) {
	var reports []domain.Report
	err := r.db.Find(&reports).Error
	return reports, err
}

func (r *reportRepository) Get(ctx context.Context, id int) (domain.Report, error) {
	var report domain.Report
	err := r.db.Where("id = ?", id).First(&report).Error
	return report, err
}

func (r *reportRepository) Create(ctx context.Context, report domain.Report) error {
	return r.db.Create(&report).Error
}

func (r *reportRepository) Update(ctx context.Context, report domain.Report, id int) error {
	return r.db.Where("id = ?", id).Updates(&report).Error
}

func (r *reportRepository) Delete(ctx context.Context, id int) error {
	return r.db.Where("id = ?", id).Delete(&domain.Report{}).Error
}
