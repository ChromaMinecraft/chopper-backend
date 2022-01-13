package main

import (
	"fmt"
	"net"
	"os"

	_reportService "github.com/ChromaMinecraft/chopper-backend/internal/core/services"
	_reportHandler "github.com/ChromaMinecraft/chopper-backend/internal/handlers/grpc"
	_reportRepo "github.com/ChromaMinecraft/chopper-backend/internal/repositories/mysql"

	"github.com/ChromaMinecraft/chopper-backend/gen/proto"

	"github.com/ChromaMinecraft/chopper-backend/utils/database"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	godotenv.Load()
}

func main() {
	addr := fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	db, err := database.InitializeDB()
	if err != nil {
		panic(err)
	}

	reportRepo := _reportRepo.NewReportRepository(db)
	reportUsecase := _reportService.NewReportService(reportRepo)
	reportHandler := _reportHandler.NewGRPCReportHandler(reportUsecase)

	grpcServer := grpc.NewServer()

	proto.RegisterReportServiceServer(grpcServer, reportHandler)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
