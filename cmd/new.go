package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "new",
		Short: "new a project",
		RunE: func(cmd *cobra.Command, args []string) error {
			path, err := exec.LookPath("kratos")
			if err != nil {
				return err
			}
		},
	}
}
