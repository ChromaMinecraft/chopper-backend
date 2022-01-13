package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ChromaMinecraft/chopper-backend/utils/convert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const mysqlDsn = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

func InitializeDB() (*gorm.DB, error) {
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
