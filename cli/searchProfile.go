package cli

import (
	"github.com/go-lens/golens/internal/pkg/lensgraphql"
	"github.com/spf13/cobra"
)

func SearchProfileCmd() *cobra.Command {
	// Build command
	cmd := &cobra.Command{
		Use:   "search",
		Short: "Search for profile",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			lensgraphql.SearchProfile(args)
		},
	}
	return cmd
}
