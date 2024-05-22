package metrics_mock_service

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	// "strconv"
	"github.com/ivanpatera99/metrics-example-app/src/domain/entities"
)

type MetricsMockService struct {
}

func (m *MetricsMockService) GetAvailableServersMetrics() ([]entities.ServerMetric, error) {
	numServers := 1
	
	serverMetrics := make([]entities.ServerMetric, numServers)
    var wg sync.WaitGroup

	// create mock data for multiple servers
    for i := 0; i < numServers; i++ {
		wg.Add(1)
        go func(i int) {
			defer wg.Done()
            mock := entities.ServerMetric{
                ID:          fmt.Sprintf("server-%d", i+1),
                CPUUsage:    rand.Float64() * 100,
                MemoryUsage: rand.Float64() * 100,
                DiskUsage:   rand.Float64() * 100,
                Timestamp:   time.Now(),
            }
            serverMetrics[i] = mock
        }(i)
    }

	wg.Wait()
    return serverMetrics, nil
}
