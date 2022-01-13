package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	_reportService "github.com/ChromaMinecraft/chopper-backend/internal/core/services"
	_reportHandler "github.com/ChromaMinecraft/chopper-backend/internal/handlers/rest"
	_reportRepo "github.com/ChromaMinecraft/chopper-backend/internal/repositories/mysql"
	"github.com/ChromaMinecraft/chopper-backend/utils/database"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	godotenv.Load()
}

func main() {
	e := echo.New()

	db, err := database.InitializeDB()
	if err != nil {
		panic(err)
	}

	reportRepo := _reportRepo.NewReportRepository(db)
	reportService := _reportService.NewReportService(reportRepo)
	reportHandler := _reportHandler.NewRestReportHandler(reportService)

	router := e.Group("/api/v1")

	reportRoute := router.Group("/report")
	reportRoute.GET("/", reportHandler.GetAll)
	reportRoute.GET("/:id", reportHandler.Get)
	reportRoute.POST("/", reportHandler.Create)
	reportRoute.PUT("/:id", reportHandler.Update)
	reportRoute.DELETE("/:id", reportHandler.Delete)

	addr := fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))

	if err := e.Start(addr); err != nil {
		e.Logger.Fatal(err)
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGTERM)
	<-sigc

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
