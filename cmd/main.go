package main

import (
	"os"

	cmd "github.com/BinacsLee/server/cmd/commands"
)

func main() {
	rootCmd := cmd.RootCmd
	rootCmd.AddCommand(
		cmd.StartCmd,
		cmd.VersionCmd,
	)
	if rootCmd.Execute() != nil {
		os.Exit(1)
	}
}
