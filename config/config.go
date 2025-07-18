/*
Copyright © 2025 Khristian Algarves khris@ueby.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

var (
	cfg  Config
	once sync.Once
)

func Load(cfgFile string) (*Config, error) {
	var loadErr error
	once.Do(func() {
		data, err := os.ReadFile(cfgFile)
		if err != nil {
			loadErr = fmt.Errorf("error reading file: %w", err)
			return
		}

		if yaml.Unmarshal(data, &cfg) != nil && json.Unmarshal(data, &cfg) != nil {
			loadErr = fmt.Errorf("error parsing configuration file. Ensure it is in valid YAML or JSON format")
			return
		}

		err = validateConfig(cfg)
		if err != nil {
			loadErr = fmt.Errorf("configuration validation failed: %w", err)
			return
		}
	})

	if loadErr != nil {
		return nil, loadErr
	}

	return &cfg, nil
}

func validateConfig(cfg Config) error {
	for _, server := range cfg.Servers {
		if server.Name == "" || server.Host == "" || server.Port <= 0 {
			return fmt.Errorf("invalid server configuration: %+v", server)
		}
	}

	db := cfg.Database
	if db.Host == "" || db.Port <= 0 || db.User == "" {
		return fmt.Errorf("invalid database configuration: %+v", db)
	}

	fmt.Println("Configuration is valid.")
	return nil
}
