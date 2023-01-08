package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func VersionCmd() *cobra.Command {
	// Hardcoded version
	currentVersion := "v0.0.1"
	// Build command
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print golens version",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(cmd.OutOrStdout(), "golens "+currentVersion)
		},
	}
	return cmd
}
