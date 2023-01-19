package main

import "github.com/go-lens/golens/cli"

func main() {
	golensCmd := cli.RootCmd()
	golensCmd.AddCommand(
		cli.VersionCmd(),
		cli.StatsCmd(),
		cli.SearchProfileCmd(),
	)
	golensCmd.Execute()
}
