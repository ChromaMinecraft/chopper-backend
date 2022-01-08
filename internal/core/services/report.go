package services

import (
	"context"

	"github.com/ChromaMinecraft/chopper-backend/internal/core/domain"
	"github.com/ChromaMinecraft/chopper-backend/internal/core/ports"
)

type reportService struct {
	Repo ports.ReportRepositorier
}

func NewReportService(repo ports.ReportRepositorier) ports.ReportServicer {
	return &reportService{
		Repo: repo,
	}
}

func (svc *reportService) GetAll(ctx context.Context) ([]domain.Report, error) {
	return svc.Repo.GetAll(ctx)
}

func (svc *reportService) Get(ctx context.Context, id int) (domain.Report, error) {
	return svc.Repo.Get(ctx, id)
}

func (svc *reportService) Create(ctx context.Context, report domain.Report) error {
	return svc.Repo.Create(ctx, report)
}

func (svc *reportService) Update(ctx context.Context, report domain.Report, id int) error {
	return svc.Repo.Update(ctx, report, id)
}

func (svc *reportService) Delete(ctx context.Context, id int) error {
	return svc.Repo.Delete(ctx, id)
}
