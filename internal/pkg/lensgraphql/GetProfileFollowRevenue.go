package lensgraphql

import (
	"context"

	graphql2 "github.com/Khan/genqlient/graphql"
)

func ListProfileFollowRevenue(args []string) {
	profileId := getFirstArg(args)
	if profileId == "" || !isValidProfileId(profileId) {
		return
	}

	ctx := context.Background()
	client := graphql2.NewClient("https://api.lens.dev/", nil)
	profileRevenue, err := profileFollowRevenue(ctx, client, ProfileFollowRevenueQueryRequest{
		// ProfileId: "0x0d", //yoginth.lens
		// ProfileId: "0x05", //stani.lens
		ProfileId: profileId,
	})
	if err != nil {
		panic(err)
	}

	printJson(profileRevenue.ProfileFollowRevenue.Revenues)
}
