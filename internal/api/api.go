package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gitlab.com/JonasEtzold/go-service-template/internal/api/router"
	"gitlab.com/JonasEtzold/go-service-template/internal/pkg/config"
	"gitlab.com/JonasEtzold/go-service-template/internal/pkg/db"
	"go.uber.org/zap"
)

func Run(logger *zap.Logger, configPath string) {
	if configPath == "" {
		configPath = "data/service.env"
	}

	config.Setup(configPath, logger)
	db.Setup(logger)
	gin.SetMode(viper.GetString("SERVER_MODE"))

	web := router.Setup(logger)
	logger.Info("go-service-template running on port " + viper.GetString("SERVER_PORT"))
	logger.Info("==================>")
	_ = web.Run(":" + viper.GetString("SERVER_PORT"))
}
