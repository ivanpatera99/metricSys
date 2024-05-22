package sql

import (
	"log"
	"database/sql"
	"github.com/ivanpatera99/metrics-example-app/src/domain/entities"
	_ "github.com/mattn/go-sqlite3"
)

type MetricsSqlAdapter struct {
	DB *sql.DB
}

func NewMetricsSqlAdapter() *MetricsSqlAdapter {
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		log.Fatal(err)
	}
	return &MetricsSqlAdapter{DB: db}
}

func (m *MetricsSqlAdapter) PostNewMetric(metrics []entities.ServerMetric) error {
	// insert metric in the database
	qry := "INSERT INTO metrics (ID, CPUUsage, MemoryUsage, DiskUsage, Timestamp) VALUES (?, ?, ?, ?, ?)"
	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}

	for _, metric := range metrics {
		_, err := tx.Exec(qry, metric.ID, metric.CPUUsage, metric.MemoryUsage, metric.DiskUsage, metric.Timestamp)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (m *MetricsSqlAdapter) GetLatestMetrics(regLimit int) ([]entities.ServerMetric, error) {
	// get all metrics from the database
	qry := "SELECT ID, CPUUsage, MemoryUsage, DiskUsage, Timestamp FROM metrics ORDER BY timestamp DESC LIMIT ?"
	stmt, err := m.DB.Prepare(qry)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(regLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var metrics []entities.ServerMetric
	for rows.Next() {
		var metric entities.ServerMetric
		err := rows.Scan(&metric.ID, &metric.CPUUsage, &metric.MemoryUsage, &metric.DiskUsage, &metric.Timestamp)
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, metric)
	}
	return metrics, nil
}