package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/JonasEtzold/go-service-template/internal/pkg/db"
)

// build information
var (
	GitCommit string
	GitTags   string
	BuildDate string
)

// HealthResponse /health response dto
type HealthResponse struct {
	Service   bool   `json:"service"`
	GitCommit string `json:"commit"`
	GitTags   string `json:"tags"`
	BuildDate string `json:"buildDate"`
	Database  bool   `json:"database"`
}

type EmptyResponse struct {
}

// Internal: no documentation
func Metrics(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, EmptyResponse{})
}

// Health /health endpoint
// Health godoc
// @Summary      Health endpoint provides information about the service operational status
// @Description  Health endpoint provides information about the service operational status
// @Produce      json
func Health(c *gin.Context) {
	res := HealthResponse{
		Service:   true,
		GitCommit: GitCommit,
		GitTags:   GitTags,
		BuildDate: BuildDate,
<% if (enableDB) { %>
		Database:  false,
<% } %>
	}
<% if (enableDB) { %>
	if db.Get() != nil {
		res.Database = true
	}
<% } %>

	c.JSON(http.StatusOK, res)
}
