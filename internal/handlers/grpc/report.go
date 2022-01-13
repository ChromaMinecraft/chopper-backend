package grpc

import (
	"context"

	"github.com/ChromaMinecraft/chopper-backend/gen/proto"
	"github.com/ChromaMinecraft/chopper-backend/internal/core/domain"
	"github.com/ChromaMinecraft/chopper-backend/internal/core/ports"
)

type grpcReportHandler struct {
	Svc ports.ReportServicer
	proto.ReportServiceServer
}

func NewGRPCReportHandler(svc ports.ReportServicer) *grpcReportHandler {
	return &grpcReportHandler{
		Svc: svc,
	}
}

func (h *grpcReportHandler) GetAll(ctx context.Context, req *proto.EmptyParam) (resp *proto.ReportResponses, err error) {
	resp = &proto.ReportResponses{
		Message: "Get all reports success",
		Report:  []*proto.Report{},
	}

	result, err := h.Svc.GetAll(ctx)
	if err != nil {
		return
	}

	for _, r := range result {
		resp.Report = append(resp.Report, &proto.Report{
			Id:          int64(r.ID),
			ReporterId:  r.ReporterID,
			ReportedId:  r.ReportedID,
			ReportId:    r.ReportID,
			Description: r.Description,
		})
	}

	return
}

func (h *grpcReportHandler) Get(ctx context.Context, req *proto.ID) (resp *proto.ReportResponse, err error) {
	resp = &proto.ReportResponse{
		Message: "Get report failed",
	}

	if req.GetId() == 0 {
		resp.Message = "Report id cannot be empty"
		return
	}

	result, err := h.Svc.Get(ctx, int(req.GetId()))
	if err != nil {
		return
	}

	resp = &proto.ReportResponse{
		Message: "Get report success",
		Report: &proto.Report{
			Id:          int64(result.ID),
			ReporterId:  result.ReporterID,
			ReportedId:  result.ReportedID,
			ReportId:    result.ReportID,
			Description: result.Description,
		},
	}

	return
}

func (h *grpcReportHandler) Create(ctx context.Context, req *proto.Report) (resp *proto.ID, err error) {
	resp = &proto.ID{
		Id: 0,
	}

	cnt := domain.Report{
		ReporterID:  req.GetReporterId(),
		ReportedID:  req.GetReportedId(),
		ReportID:    req.GetReportId(),
		Description: req.GetDescription(),
	}

	if err = h.Svc.Create(ctx, cnt); err != nil {
		return
	}

	return
}

func (h *grpcReportHandler) Update(ctx context.Context, req *proto.Report) (resp *proto.ID, err error) {
	resp = &proto.ID{
		Id: 0,
	}

	cnt := domain.Report{
		ID:          int(req.GetId()),
		ReporterID:  req.GetReporterId(),
		ReportedID:  req.GetReportedId(),
		ReportID:    req.GetReportId(),
		Description: req.GetDescription(),
	}

	if err = h.Svc.Update(ctx, cnt, cnt.ID); err != nil {
		return
	}

	resp.Id = int64(cnt.ID)

	return
}

func (h *grpcReportHandler) Delete(ctx context.Context, req *proto.ID) (resp *proto.ID, err error) {
	resp = &proto.ID{
		Id: 0,
	}

	if err = h.Svc.Delete(ctx, int(req.GetId())); err != nil {
		return
	}

	resp.Id = int64(req.GetId())

	return
}
