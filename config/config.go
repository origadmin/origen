package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/origadmin/toolkits/codec"
)

const Application = `Origen`

// Config is the configuration of admin-cli.
type Config struct {
	Application string    `json:"application" yaml:"application" toml:"application"`    // app_name
	PackageName string    `json:"package_name" yaml:"package_name" toml:"package_name"` // package_name
	Author      string    `json:"author" yaml:"author" toml:"author"`                   // author
	WirePath    string    `json:"wire_path" yaml:"wire_path" toml:"wire_path"`          // wire_path
	SwaggerPath string    `json:"swagger_path" yaml:"swagger_path" toml:"swagger_path"` // swagger_path
	Fontend     Fontend   `json:"fontend" yaml:"fontend" toml:"fontend"`                // fontend
	Backend     Backend   `json:"backend" yaml:"backend" toml:"backend"`                // backend
	Database    Database  `json:"database" yaml:"database" toml:"database"`             // database
	Extension   Extension `json:"extension" yaml:"extension" toml:"extension"`          // extension
	Description string    `json:"description" yaml:"description" toml:"description"`    // description
	Version     string    `json:"version" yaml:"version" toml:"version"`                // version
}

// Save saves config to file
func Save(path string, v any) error {
	dir, _ := filepath.Split(path)
	// check whether the directory exists
	stat, err := os.Stat(dir)
	if err != nil {
		// if the directory does not exist create it
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}
	if err == nil && !stat.IsDir() {
		// if the path exists but is not a directory, an error is returned
		return fmt.Errorf("%s is not a directory", dir)
	}

	return codec.EncodeFile(path, v)
}

// Load loads config from file
func Load(path string, v any) error {
	stat, err := os.Stat(path)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return fmt.Errorf("%s is a directory", path)
	}
	return codec.DecodeFile(path, v)
}

var (
	config Config
	once   sync.Once
)

// C is a singleton of Config
func C() Config {
	return config
}
