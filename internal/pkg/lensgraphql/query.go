package lensgraphql

import (
	"context"

	"github.com/modood/table"
	"github.com/shurcooL/graphql"
)

type Response struct {
	GlobalProtocolStats GlobalProtocolStats `json:"globalProtocolStats"`
}

func ListStats() {
	// set up GraphQL query
	var q GraphQLQuery
	q.Query = GlobalProtocolStatsQuery
	// Se up GraphQL Client
	client := graphql.NewClient("https://api.lens.dev/", nil)

	// Execute the query
	var r Response
	if err := client.Query(context.Background(), &r, nil); err != nil {
		panic(err)
	}

	// Print the response as table
	outputResp := []GlobalProtocolStats{r.GlobalProtocolStats}
	table.Output(outputResp)
}
