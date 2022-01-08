package ports

import (
	"context"

	"github.com/ChromaMinecraft/chopper-backend/internal/core/domain"
)

type ReportServicer interface {
	GetAll(ctx context.Context) ([]domain.Report, error)
	Get(ctx context.Context, id int) (domain.Report, error)
	Create(ctx context.Context, report domain.Report) error
	Update(ctx context.Context, report domain.Report, id int) error
	Delete(ctx context.Context, id int) error
}
