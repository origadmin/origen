// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package config is the config package for origen
package config

type Application struct {
	Type         string // 例如: "web", "mobile", "desktop"
	TemplatePath string
	OutputPath   string
}
