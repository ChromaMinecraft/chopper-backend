package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ChromaMinecraft/chopper-backend/internal/core/domain"
	"github.com/ChromaMinecraft/chopper-backend/internal/core/ports"
	"github.com/labstack/echo/v4"
)

type restReportHandler struct {
	Svc ports.ReportServicer
}

func NewRestReportHandler(svc ports.ReportServicer) *restReportHandler {
	return &restReportHandler{
		Svc: svc,
	}
}

func (h *restReportHandler) GetAll(c echo.Context) error {
	ctx := context.Background()

	result, err := h.Svc.GetAll(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, BuildRESTResponse(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, BuildRESTResponse("Get all reports success", result))
}

func (h *restReportHandler) Get(c echo.Context) error {
	ctx := context.Background()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, BuildRESTResponse(err.Error(), nil))
	}

	result, err := h.Svc.Get(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, BuildRESTResponse(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, BuildRESTResponse("Get report success", result))
}

func (h *restReportHandler) Create(c echo.Context) error {
	ctx := context.Background()

	var report domain.Report
	if err := c.Bind(&report); err != nil {
		return c.JSON(http.StatusBadRequest, BuildRESTResponse(err.Error(), nil))
	}

	if err := h.Svc.Create(ctx, report); err != nil {
		return c.JSON(http.StatusInternalServerError, BuildRESTResponse(err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, BuildRESTResponse("Create report success", nil))
}

func (h *restReportHandler) Update(c echo.Context) error {
	ctx := context.Background()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, BuildRESTResponse(err.Error(), nil))
	}

	var report domain.Report
	if err := c.Bind(&report); err != nil {
		return c.JSON(http.StatusBadRequest, BuildRESTResponse(err.Error(), nil))
	}

	if err := h.Svc.Update(ctx, report, id); err != nil {
		return c.JSON(http.StatusInternalServerError, BuildRESTResponse(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, BuildRESTResponse("Update report success", nil))
}

func (h *restReportHandler) Delete(c echo.Context) error {
	ctx := context.Background()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, BuildRESTResponse(err.Error(), nil))
	}

	if err := h.Svc.Delete(ctx, id); err != nil {
		return c.JSON(http.StatusInternalServerError, BuildRESTResponse(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, BuildRESTResponse("Delete report success", nil))
}
