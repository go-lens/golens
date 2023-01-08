package cli

import (
	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	// build command
	cmd := &cobra.Command{
		Use:   "golens",
		Short: "A brief description of your application",
		Long:  "A Go wrapper for Lens Protocol Graphql queries",
	}

	return cmd
}
