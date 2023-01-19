package cli

import (
	"github.com/go-lens/golens/internal/pkg/lensgraphql"
	"github.com/spf13/cobra"
)

func ListProfileFollowRevenue() *cobra.Command {
	// Build command
	cmd := &cobra.Command{
		Use:   "revenue",
		Short: "Get revenue based on profile followers",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			lensgraphql.ListProfileFollowRevenue(args)
		},
	}
	return cmd
}
