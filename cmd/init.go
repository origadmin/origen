// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package cmd defines a CLI command to start a server with various flags and options, including the
// ability to run as a daemon.
// It includes functions to create and manage the command, as well as the logic to run the server.
// It also includes a function to create a new server instance and start it.
package cmd

import (
	"github.com/spf13/cobra"
)

func InitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "init the project",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
