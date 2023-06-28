package router

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gitlab.com/JonasEtzold/go-service-template/internal/api/controllers"
	"gitlab.com/JonasEtzold/go-service-template/internal/api/middlewares"
	"go.uber.org/zap"
)

func Setup(logger *zap.Logger) *gin.Engine {
	app := gin.New()
	// Middlewares
	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	app.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	// Logs all panic to error log
	//   - stack means whether output the stack info.
	app.Use(ginzap.RecoveryWithZap(logger, true))
	app.Use(middlewares.CORS())
	app.NoRoute(middlewares.NoRouteHandler())

	// Routes
	// ================== Info routes
	app.GET("/health", controllers.Health)
	app.GET("/metrics", controllers.Metrics)
	// ================== Docs Routes
	app.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// ================== Login Routes
	app.POST("/api/v1/login", controllers.Login)
	// ================== Authenticated Routes
	authorized := app.Group("/")
	authorized.Use(middlewares.AuthRequired())
	{
		// ================== Example Routes
		authorized.GET("/api/v1/example/:id", controllers.GetExampleById)
		authorized.GET("/api/v1/example", controllers.GetExample)
		authorized.POST("/api/v1/example", controllers.CreateExample)
		authorized.PUT("/api/v1/example/:id", controllers.UpdateExample)
		authorized.DELETE("/api/v1/example/:id", controllers.DeleteExample)
	}
	return app
}
