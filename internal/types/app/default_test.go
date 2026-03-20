package app

import (
	"testing"
)

func TestServerHostDefault(t *testing.T) {
	tests := []struct {
		name         string
		inputHost    string
		expectedHost string
	}{
		{
			name:         "empty host should default to 0.0.0.0",
			inputHost:    "",
			expectedHost: "0.0.0.0",
		},
		{
			name:         "custom host should be preserved",
			inputHost:    "127.0.0.1",
			expectedHost: "127.0.0.1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &Config{
				ServerHost: tt.inputHost,
				ServerPort: 5002,
				ServerKey:  "test-key",
			}
			config.SetDefault()

			if config.ServerHost != tt.expectedHost {
				t.Errorf("expected ServerHost %s, got %s", tt.expectedHost, config.ServerHost)
			}
		})
	}
}

func TestKingbaseESDBTypeDefault(t *testing.T) {
	tests := []struct {
		name                  string
		inputDBType           string
		expectedDBType        string
		expectedDefaultDBName string
	}{
		{
			name:                  "kingbase should be aliased to postgresql",
			inputDBType:           DB_TYPE_KINGBASE,
			expectedDBType:        DB_TYPE_POSTGRESQL,
			expectedDefaultDBName: "postgres",
		},
		{
			name:                  "postgresql should remain postgresql",
			inputDBType:           DB_TYPE_POSTGRESQL,
			expectedDBType:        DB_TYPE_POSTGRESQL,
			expectedDefaultDBName: "postgres",
		},
		{
			name:                  "pgbouncer should remain pgbouncer",
			inputDBType:           DB_TYPE_PG_BOUNCER,
			expectedDBType:        DB_TYPE_PG_BOUNCER,
			expectedDefaultDBName: "postgres",
		},
		{
			name:                  "mysql should remain mysql",
			inputDBType:           DB_TYPE_MYSQL,
			expectedDBType:        DB_TYPE_MYSQL,
			expectedDefaultDBName: "mysql",
		},
		{
			name:                  "oceanbase should be aliased to mysql",
			inputDBType:           DB_TYPE_OCEANBASE,
			expectedDBType:        DB_TYPE_MYSQL,
			expectedDefaultDBName: "mysql",
		},
		{
			name:                  "seekdb should be aliased to mysql",
			inputDBType:           DB_TYPE_SEEKDB,
			expectedDBType:        DB_TYPE_MYSQL,
			expectedDefaultDBName: "mysql",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &Config{
				DBType: tt.inputDBType,
			}
			config.SetDefault()

			if config.DBType != tt.expectedDBType {
				t.Errorf("expected DBType %s, got %s", tt.expectedDBType, config.DBType)
			}
			if config.DBDefaultDatabase != tt.expectedDefaultDBName {
				t.Errorf("expected DBDefaultDatabase %s, got %s", tt.expectedDefaultDBName, config.DBDefaultDatabase)
			}
		})
	}
}
