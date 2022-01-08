package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_reportService "github.com/ChromaMinecraft/chopper-backend/internal/core/services"
	_reportHandler "github.com/ChromaMinecraft/chopper-backend/internal/handlers/rest"
	_reportRepo "github.com/ChromaMinecraft/chopper-backend/internal/repositories/mysql"
	"github.com/ChromaMinecraft/chopper-backend/utils/convert"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	mysqlDsn = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

func init() {
	godotenv.Load()
}

func main() {
	e := echo.New()

	db, err := initializeDB()
	if err != nil {
		panic(err)
	}

	reportRepo := _reportRepo.NewReportRepository(db)
	reportService := _reportService.NewReportService(reportRepo)
	reportHandler := _reportHandler.NewRestHandler(reportService)

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

func initializeDB() (*gorm.DB, error) {
	gormConf := &gorm.Config{}

	dbDebug := convert.StringToBool(os.Getenv("DB_DEBUG"), false)
	if dbDebug {
		gormConf.Logger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: 2 * time.Second,
				LogLevel:      logger.Silent,
				Colorful:      true,
			},
		)
	}

	dsn := fmt.Sprintf(mysqlDsn, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), gormConf)
	if err != nil {
		return nil, err
	}

	if dbDebug {
		return db.Debug(), nil
	}

	return db, nil
}
