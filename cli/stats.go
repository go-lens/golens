package cli

import (
	"github.com/0xkeivin/golens/internal/pkg/lensgraphql"
	"github.com/spf13/cobra"
)

func StatsCmd() *cobra.Command {
	// Build command
	cmd := &cobra.Command{
		Use:   "stats",
		Short: "Prints global lens protocol stats",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			lensgraphql.ListStats()
		},
	}
	return cmd
}
