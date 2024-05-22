CREATE TABLE IF NOT EXISTS metrics (
    ID VARCHAR(255),
    CPUUsage DECIMAL(5,2),
    MemoryUsage DECIMAL(5,2),
    DiskUsage DECIMAL(5,2),
    Timestamp TIMESTAMP,
    PRIMARY KEY (ID, Timestamp)
);