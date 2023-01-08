package lensgraphql

type GraphQLQuery struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}
type GlobalProtocolStats struct {
	TotalProfiles      int `json:"totalProfiles"`
	TotalBurntProfiles int `json:"totalBurntProfiles"`
	TotalPosts         int `json:"totalPosts"`
	TotalMirrors       int `json:"totalMirrors"`
	TotalComments      int `json:"totalComments"`
	TotalCollects      int `json:"totalCollects"`
	TotalFollows       int `json:"totalFollows"`
	// TotalRevenue       []int `json:"totalRevenue"`
}
