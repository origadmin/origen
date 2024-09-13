// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package config is the config package for origen
package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/origadmin/toolkits/codec"
)

const (
	Project     = "origin"
	PackageName = "application"
)

// Config is the configuration of admin-cli.
type Config struct {
	Project     string             `json:"project" default:"origin"`
	PackageName string             `json:"package_name" default:"application"`
	Author      string             `json:"author" default:"Origen"`
	WirePath    string             `json:"wire_path"`
	SwaggerPath string             `json:"swagger_path"`
	Description string             `json:"description"`
	Application Application        `json:"application" yaml:"application" toml:"application"` // application
	Fontend     Fontend            `json:"fontend" yaml:"fontend" toml:"fontend"`             // fontend
	Backend     Backend            `json:"backend" yaml:"backend" toml:"backend"`             // backend
	Database    Database           `json:"database" yaml:"database" toml:"database"`          // database
	Extension   Extension          `json:"extension" yaml:"extension" toml:"extension"`       // extension
	Services    map[string]Service `json:"services" yaml:"services" toml:"services"`          // services
	Version     string             `json:"version"`
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

// LoadGlobal loads config from file
func LoadGlobal(path string) {
	once.Do(func() {
		err := Load(path, &config)
		if err != nil {
			panic(err)
		}
	})
}
