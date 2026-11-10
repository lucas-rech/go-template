package main

import (
	"context"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "github.com/lucas-rech/go-template/api/docs"
	"github.com/lucas-rech/go-template/internal/config"
	"github.com/lucas-rech/go-template/observability/logger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	env, err := config.Load()
	if err != nil {
		slog.Error("Error loading config", err)
	}

	logger.SetupDefault(env.LogFormat, env.LogLevel)
	ctx := logger.WithTraceID(context.Background(), "main-startup")
	slog.InfoContext(ctx, "Starting application", "name", env.AppName, "api_version", config.Version)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "pong"})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
