package rest

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ivanpatera99/metrics-app/src/domain/ports"
	"github.com/ivanpatera99/metrics-app/src/domain/usecases"
)

func App(metricsRepo ports.MetricsRepo, metricsService ports.MetricsService) *gin.Engine {
	metricsUseCase := usecases.MetricsUseCase{
		MetricsRepo:           metricsRepo,
		ServerMetricsProvider: metricsService,
	}

	r := gin.Default()

	// enable access via /metrics endpoint
	r.GET("/metrics", func(c *gin.Context) {
		// call usecase
		limitStr := c.Query("limit")
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid limit query parameter"})
			return
		}
		metrics, err := metricsUseCase.GetLatestMetrics(limit)
		if err != nil {
			c.JSON(500, gin.H{"error": "internal server error"})
			return
		}
		c.JSON(200, gin.H{"metrics": metrics})
		// Use the "limit" value in your use case logic
	})

	// boot up go routine for generating metrics
	go metricsUseCase.GenerateMetrics()

	return r
}
