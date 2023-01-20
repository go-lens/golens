package exploreProfile

import (
	"context"

	graphql2 "github.com/Khan/genqlient/graphql"
	"github.com/go-lens/golens/internal/pkg/lensgraphql"
)

func ExploreProfiles(sortCriteriaInput string, cursor string) {
	// handle sortCriteriaInput
	var sortCriteria ProfileSortCriteria
	sortCriteriaMap := map[string]ProfileSortCriteria{
		"1": "CREATED_ON",
		"2": "MOST_FOLLOWERS",
		"3": "LATEST_CREATED",
		"4": "MOST_POSTS",
		"5": "MOST_COMMENTS",
		"6": "MOST_MIRRORS",
		"7": "MOST_PUBLICATION",
		"8": "MOST_COLLECTS",
	}
	if sortCriteriaInput == "" {
		sortCriteria = sortCriteriaMap["1"]
	} else {
		sortCriteria = sortCriteriaMap[sortCriteriaInput]
	}

	cursor = "{\"offset\":" + cursor + "}"
	// fmt.Printf("cursor: %s\n", cursor)
	// fmt.Printf("sortCriteriaInput: %s\n", sortCriteriaInput)
	// fmt.Printf("sortCriteria: %s\n", sortCriteria)
	ctx := context.Background()
	client := graphql2.NewClient("https://api.lens.dev/", nil)
	profiles, err := exploreProfiles(ctx, client,
		cursor,
		sortCriteria,
	)

	if err != nil {
		panic(err)
	}
	lensgraphql.PrintJson(profiles)

}
