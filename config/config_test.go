package config

import "testing"

func TestValidConfig(t *testing.T) {
	cfg := Config{
		Servers: []ServerConfig{
			{
				Name:     "web",
				Host:     "localhost",
				Port:     8080,
				Replicas: 2,
			},
		},
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     5432,
			User:     "admin",
			Password: "secret",
		},
	}
	if cfg.Servers[0].Name == "" {
		t.Error("Expected server name to be set, but it is empty")
	}
	if cfg.Database.User != "admin" {
		t.Errorf("Expected user 'admin', but it %s", cfg.Database.User)
	}
}
