package usecases

import (
	"fmt"
	"time"

	"github.com/ivanpatera99/metrics-example-app/src/domain/entities"
	"github.com/ivanpatera99/metrics-example-app/src/domain/ports"
)

type MetricsUseCase struct {
	MetricsRepo           ports.MetricsRepo
	ServerMetricsProvider ports.MetricsService
}

func (m *MetricsUseCase) GetLatestMetrics(regLimit int) ([]entities.ServerMetric, error) {
	// get latest metrics stored in the database
	metrics, err := m.MetricsRepo.GetLatestMetrics(regLimit)
	if err != nil {
		return nil, err
	}
	return metrics, nil
}

func (m *MetricsUseCase) GenerateMetrics() error {
	for {
		// get/generate metrics from the server
		metrics, err := m.ServerMetricsProvider.GetAvailableServersMetrics()
		fmt.Println("new metric: ", metrics)
		if err != nil {
			return fmt.Errorf("error getting server metrics: %w", err)
		}
		// store them in the database
		err = m.MetricsRepo.PostNewMetric(metrics)
		if err != nil {
			return fmt.Errorf("error posting new metric: %w", err)
		}
		time.Sleep(1 * time.Second)
	}
}
