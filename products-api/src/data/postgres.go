package data

import (
	"fmt"
	"time"

	"github.com/corey888773/ztp-shopping-cart/products-api/src/data/products"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresConfig holds configuration for PostgreSQL database
type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// PostgresConnector manages the database connection
type PostgresConnector struct {
	DB *gorm.DB
}

// NewPostgresConnector creates a new database connection
func NewPostgresConnector(config PostgresConfig) (*PostgresConnector, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.Password, config.DBName, config.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Register models

	// auto migrate the database
	err = db.AutoMigrate(&products.Product{})
	err = db.AutoMigrate(&products.ProductReservation{})
	if err != nil {
		return nil, err
	}

	return &PostgresConnector{
		DB: db,
	}, nil
}

// Close closes the database connection
func (pc *PostgresConnector) Close() error {
	sqlDB, err := pc.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
