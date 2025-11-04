package testutil

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"shoplite/config"
	"shoplite/internal/models"
)

type testConfig struct {
	ServerPort string `yaml:"serverPort"`
	Database   struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		SSLMode  string `yaml:"sslmode"`
		Schema   string `yaml:"schema"`
	} `yaml:"database"`
}

// LoadTestConfig reads config.test.yaml from repo root.
func LoadTestConfig(t *testing.T) *config.Config {
	t.Helper()
	_, currentFilePath, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(currentFilePath)
	repoRoot := filepath.Dir(filepath.Dir(currentDir))
	absolutePath := filepath.Join(repoRoot, "config.test.yaml")
	b, err := os.ReadFile(absolutePath)
	if err != nil {
		t.Fatalf("failed to read config.test.yaml: %v", err)
	}
	var tc testConfig
	if err := yaml.Unmarshal(b, &tc); err != nil {
		t.Fatalf("failed to parse config.test.yaml: %v", err)
	}
	cfg := &config.Config{
		ServerPort: tc.ServerPort,
		DBHost:     tc.Database.Host,
		DBPort:     tc.Database.Port,
		DBUser:     tc.Database.User,
		DBPassword: tc.Database.Password,
		DBName:     tc.Database.Name,
		DBSSLMode:  tc.Database.SSLMode,
	}
	return cfg
}

// NewTestDB opens a connection to the test DB and runs migrations.
func NewTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	cfg := LoadTestConfig(t)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect test database: %v", err)
	}
	// AutoMigrate before each package tests
	if err := db.AutoMigrate(&models.Customer{}, &models.Product{}, &models.Order{}, &models.OrderItem{}); err != nil {
		t.Fatalf("failed to migrate test database: %v", err)
	}
	return db
}

// TruncateAll clears tables between tests.
func TruncateAll(t *testing.T, db *gorm.DB) {
	t.Helper()
	// Order matters due to FKs
	if err := db.Exec("TRUNCATE TABLE order_items, orders, products, customers RESTART IDENTITY CASCADE").Error; err != nil {
		t.Fatalf("failed to truncate tables: %v", err)
	}
}
