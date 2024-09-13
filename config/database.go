// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package config is the config package for origen
package config

type Database struct {
	Enable   bool   `json:"enable" yaml:"enable" toml:"enable"`
	Dialect  string // 例如: "mysql", "postgresql", "sqlite"
	Host     string
	Port     int
	Username string
	Password string
	Database string
}
