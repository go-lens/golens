package main

import "github.com/0xkeivin/golens/cli"

func main() {
	golensCmd := cli.RootCmd()
	golensCmd.AddCommand(
		cli.VersionCmd(),
		cli.StatsCmd(),
	)
	golensCmd.Execute()
}
