package lensgraphql

import (
	"context"

	graphql2 "github.com/Khan/genqlient/graphql"
)

func GetProfileFollowRevenue(args []string) {
	// if len(args) < 1 || len(args[0]) < 5 {
	// 	fmt.Println("Please provide a valid Lens handle (e.g. stani.lens)")
	// 	return
	// }
	// handle := args[0]
	// if !strings.HasSuffix(handle, ".lens") {
	// 	handle = handle + ".lens"
	// }
	// check if profileId is valid hex
	if !isValidProfileId(args[0]) {
		return
	}
	profileId := args[0]

	ctx := context.Background()
	client := graphql2.NewClient("https://api.lens.dev/", nil)
	profileRevenue, err := profileFollowRevenue(ctx, client, ProfileFollowRevenueQueryRequest{
		// ProfileId: "0x0d", //yoginth.lens
		// ProfileId: "0x05", //stani.lens
		// ProfileId: "0x05", //stani.lens
		ProfileId: profileId,
	})
	if err != nil {
		panic(err)
	}

	// if profileRevenue == "" {
	// 	fmt.Printf("Profile not found: %q\n", handle)
	// 	return
	// }
	printJson(profileRevenue)
}
