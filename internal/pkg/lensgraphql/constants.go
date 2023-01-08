package lensgraphql

const GlobalProtocolStatsQuery = `
query GlobalProtocolStats {
	globalProtocolStats(request: null) {
	  totalProfiles
	  totalBurntProfiles
	  totalPosts
	  totalMirrors
	  totalComments
	  totalCollects
	  totalFollows
	}
  }
`
