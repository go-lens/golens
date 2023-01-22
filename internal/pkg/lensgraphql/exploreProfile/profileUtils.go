package exploreProfile

type FlattenProfile struct {
	Handle string `json:"handle"`
	ID     string `json:"id"`
	Ens    string `json:"ens"`
	// OnChainIdentity map[string]interface{} `json:"onChainIdentity"`
	OwnedBy           string `json:"ownedBy"`
	TotalCollects     int    `json:"totalCollects"`
	TotalComments     int    `json:"totalComments"`
	TotalFollowers    int    `json:"totalFollowers"`
	TotalFollowing    int    `json:"totalFollowing"`
	TotalMirrors      int    `json:"totalMirrors"`
	TotalPosts        int    `json:"totalPosts"`
	TotalPublications int    `json:"totalPublications"`
	Metadata          string `json:"metadata"`
}

// create json from *exploreProfilesResponse
func flattenProfileJson(data *exploreProfilesResponse) []FlattenProfile {
	var result []FlattenProfile
	for _, value := range data.ExploreProfiles.Items {
		result = append(result, FlattenProfile{
			Handle:            value.Handle,
			ID:                value.Id,
			Ens:               value.OnChainIdentity.Ens.Name,
			OwnedBy:           value.OwnedBy,
			TotalCollects:     value.Stats.TotalCollects,
			TotalComments:     value.Stats.TotalComments,
			TotalFollowers:    value.Stats.TotalFollowers,
			TotalFollowing:    value.Stats.TotalFollowing,
			TotalMirrors:      value.Stats.TotalMirrors,
			TotalPosts:        value.Stats.TotalPosts,
			TotalPublications: value.Stats.TotalPublications,
			Metadata:          value.Metadata,
		})
	}
	return result
}

// func flattenProfileJson(data map[string]interface{}, prefix string) map[string]interface{} {
// 	result := make(map[string]interface{})
// 	for key, value := range data {
// 		newKey := prefix + "." + key
// 		switch value.(type) {
// 		case map[string]interface{}:
// 			for k, v := range flatten(value.(map[string]interface{}), newKey) {
// 				result[k] = v
// 			}
// 		default:
// 			result[newKey] = value
// 		}
// 	}
// 	return result
// }
