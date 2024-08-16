// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package config is the config package for origen
package config_test

import (
	"testing"

	"github.com/origadmin/toolkits/codec"

	"github.com/origadmin/origen/config"
)

func TestC(t *testing.T) {
	config := config.Config{
		Application: "Origen",
		PackageName: "origen",
		Author:      "OrigAdmin",
		WirePath:    "wire",
		SwaggerPath: "swag",
		Fontend: config.Fontend{
			Enable: false,
		},
		Backend: config.Backend{
			Enable: true,
		},
		Database: config.Database{
			Enable: false,
		},
		Extension: config.Extension{
			Enable: false,
		},
		Services: map[string]config.Service{
			"user": {
				Enable: true,
			},
		},
		Description: "Origen is a CLI scaffolding for quickly setting up a project.",
		Version:     "v0.0.0",
	}

	err := codec.EncodeTOMLFile("config.toml", config)
	if err != nil {
		t.Fatal(err)
	}
}
