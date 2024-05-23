package rest_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ivanpatera99/metrics-app/src/app/rest"
	"github.com/ivanpatera99/metrics-app/src/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMetricsRepo struct {
	mock.Mock
}

func (m *MockMetricsRepo) PostNewMetric(metrics []entities.ServerMetric) error {
	args := m.Called(metrics)
	return args.Error(0)
}
func (m *MockMetricsRepo) GetLatestMetrics(regLimit int) ([]entities.ServerMetric, error) {
	args := m.Called(regLimit)
	return args.Get(0).([]entities.ServerMetric), args.Error(1)
}

type MockMetricsService struct {
	mock.Mock
}

func (m *MockMetricsService) GetAvailableServersMetrics() ([]entities.ServerMetric, error) {
	args := m.Called()
	return args.Get(0).([]entities.ServerMetric), args.Error(1)
}

func setup() (*MockMetricsRepo, *MockMetricsService) {
	metricsRepo := new(MockMetricsRepo)
	metricsService := new(MockMetricsService)
	metricsService.On("GetAvailableServersMetrics").Return([]entities.ServerMetric{
		{
			ID:        "server-1",
			CPUUsage:  0.5,
			MemoryUsage:  0.5,
			DiskUsage: 0.5,
			Timestamp: time.Now(),
		},
	}, nil)
	metricsRepo.On("PostNewMetric", mock.Anything).Return(nil)
	return metricsRepo, metricsService
}

func TestAppRespondsWithMetrics(t *testing.T) {
	// Mock sql and metrics service
	metricsRepo, metricsService := setup()
	metricsRepo.On("GetLatestMetrics", 1).Return([]entities.ServerMetric{
		{
			ID:        "server-1",
			CPUUsage:  0.5,
			MemoryUsage:  0.5,
			DiskUsage: 0.5,
			Timestamp: time.Now(),
		},
	}, nil)
	// Given
	router := rest.App(metricsRepo, metricsService)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/metrics?limit=1", nil)
	router.ServeHTTP(w, req)
	// Then
	assert.Equal(t, 200, w.Code)
}

func TestAppHandlesUnvalidLimitShouldReturn400(t *testing.T) {
	// Mock sql and metrics service
	metricsRepo, metricsService := setup()
	// Given
	router := rest.App(metricsRepo, metricsService)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/metrics?limit=invalid", nil)
	router.ServeHTTP(w, req)
	// Then
	assert.Equal(t, 400, w.Code)
}

func TestAppFailsToGetLatestMetricsShouldReturn500(t *testing.T) {
	// Mock sql and metrics service
	metricsRepo, metricsService := setup()
	metricsRepo.On("GetLatestMetrics", 1).Return([]entities.ServerMetric{}, errors.New("error"))
	// Given
	router := rest.App(metricsRepo, metricsService)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/metrics?limit=1", nil)
	router.ServeHTTP(w, req)
	// Then
	assert.Equal(t, 500, w.Code)
}
