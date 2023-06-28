package main

import (
	_ "gitlab.com/JonasEtzold/go-service-template/docs"
	"gitlab.com/JonasEtzold/go-service-template/internal/api"
	"go.uber.org/zap"
)

// @Golang       Go Service Template
// @version      1.0
// @title        Go Service starter
// @description  A template for a Go backend API service

// @contact.name   Jonas Etzold
// @contact.email  jonas.etzold@beeyou.de

// @license.name  WTFPL
// @license.url   http://www.wtfpl.net/txt/copying/

// @BasePath  /

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	api.Run(logger, "")
}
