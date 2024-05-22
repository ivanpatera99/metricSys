package rest

import (
	"strconv"
	"github.com/gin-gonic/gin"
	metricsMockService "github.com/ivanpatera99/metrics-example-app/src/adapters/metrics/mock"
	sqlAdapter "github.com/ivanpatera99/metrics-example-app/src/adapters/sql"
	"github.com/ivanpatera99/metrics-example-app/src/domain/usecases"
)

func App() {
	metricsSqlAdapter := sqlAdapter.NewMetricsSqlAdapter()
	metricsMockService := metricsMockService.MetricsMockService{}
	metricsUseCase := usecases.MetricsUseCase{
		MetricsRepo:          metricsSqlAdapter,
		ServerMetricsProvider: &metricsMockService,
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

	// boot up the gin server
	r.Run(":8080")
}

