package ports

import "github.com/ivanpatera99/metrics-example-app/src/domain/entities"

type MetricsRepo interface {
	PostNewMetric(metrics []entities.ServerMetric) error
	GetLatestMetrics(regLimit int) ([]entities.ServerMetric, error)
}

type MetricsService interface {
	GetAvailableServersMetrics() ([]entities.ServerMetric, error)
}