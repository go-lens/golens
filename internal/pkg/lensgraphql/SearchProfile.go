package lensgraphql

import (
	"context"
	"fmt"
	"strings"

	graphql2 "github.com/Khan/genqlient/graphql"
)

func SearchProfile(args []string) {
	if len(args) < 1 || len(args[0]) < 5 {
		fmt.Println("Please provide a valid Lens handle (e.g. stani.lens)")
		return
	}
	handle := args[0]
	if !strings.HasSuffix(handle, ".lens") {
		handle = handle + ".lens"
	}

	ctx := context.Background()
	client := graphql2.NewClient("https://api.lens.dev/", nil)
	profileValue, err := profile(ctx, client, SingleProfileQueryRequest{
		// Handle: "stani.lens",
		// ProfileId: "0x05",
		Handle: handle,
	})
	if err != nil {
		panic(err)
	}

	if profileValue.Profile.Handle == "" {
		fmt.Printf("Profile not found: %q\n", handle)
		return
	}
	printJson(profileValue.Profile)
}
