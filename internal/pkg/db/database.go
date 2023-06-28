package db

import (
	"time"

	"github.com/spf13/viper"
	"gitlab.com/JonasEtzold/go-service-template/internal/pkg/models/example"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

var (
	DB  *gorm.DB
	err error
)

type Database struct {
	*gorm.DB
}

// SetupDB opens a database and saves the reference to `Database` struct.
func Setup(logger *zap.Logger) {
	var db = DB
	dbLogger := zapgorm2.New(logger)
	dbLogger.SetAsDefault()

	driver := viper.GetString("database_driver")
	database := viper.GetString("database_name")
	username := viper.GetString("database_username")
	password := viper.GetString("database_password")
	host := viper.GetString("database_host")
	port := viper.GetString("database_port")

	logger.Info("Database: using " + driver + " driver for connecting.")
	if driver == "sqlite" { // SQLITE
		db, err = gorm.Open(sqlite.Open("./"+database+".db"), &gorm.Config{
			Logger: dbLogger,
		})
		if err != nil {
			logger.Error("db err: ", zap.Error(err))
		}
	} else if driver == "postgres" { // POSTGRES
		dsn := "host=" + host + " port=" + port + " user=" + username + " dbname=" + database + "  sslmode=disable password=" + password
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: dbLogger,
		})
		if err != nil {
			logger.Error("db err: ", zap.Error(err))
		}
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("db err: ", zap.Error(err))
	}

	sqlDB.SetMaxIdleConns(viper.GetInt("database_idle_conns"))
	sqlDB.SetMaxOpenConns(viper.GetInt("database_open_conns"))
	sqlDB.SetConnMaxLifetime(time.Duration(viper.GetInt("database_max_lifetime")) * time.Second)

	DB = db
	migration(logger)
}

// Auto migrate project models
func migration(logger *zap.Logger) {
	logger.Info("Database: creating service model tables")
	DB.AutoMigrate(&example.Example{})
}

func Get() *gorm.DB {
	return DB
}
