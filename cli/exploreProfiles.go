package cli

import (
	"github.com/go-lens/golens/internal/pkg/lensgraphql/exploreProfile"
	"github.com/spf13/cobra"
)

func ExploreProfilesCmd() *cobra.Command {
	var cursor string
	// Build command
	cmd := &cobra.Command{
		Use:   "explore",
		Short: "Explore profiles based on criteria",
		Long: `Explore profiles based on criteria:
	1 - CREATED_ON (default)
	2 - MOST_FOLLOWERS
	3 - LATEST_CREATED
	4 - MOST_POSTS
	5 - MOST_COMMENTS
	6 - MOST_MIRRORS
	7 - MOST_PUBLICATION
	8 - MOST_COLLECTS
`,
		Args:      cobra.MaximumNArgs(1),
		ValidArgs: []string{"1", "2", "3", "4", "5", "6", "7", "8"},
		Run: func(cmd *cobra.Command, args []string) {
			var sortCriteriaInput string
			if len(args) < 1 {
				sortCriteriaInput = "1"
			} else {
				sortCriteriaInput = args[0]
			}
			exploreProfile.ExploreProfiles(sortCriteriaInput, cursor)
		},
	}
	cmd.Flags().StringVarP(&cursor, "cursor", "c", "0", "Cursor to start from")
	return cmd
}
