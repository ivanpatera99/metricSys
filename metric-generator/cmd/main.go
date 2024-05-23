package main

import (
	"github.com/ivanpatera99/metrics-app/src/app/rest"
	metricsMockService "github.com/ivanpatera99/metrics-app/src/adapters/metrics/mock"
	sqlAdapter "github.com/ivanpatera99/metrics-app/src/adapters/sql"
)

func main() {
	rest.App(sqlAdapter.NewMetricsSqlAdapter(),  &metricsMockService.MetricsMockService{}).Run(":8080")
}
