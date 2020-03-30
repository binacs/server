package main

import (
	//"github.com/spf13/cobra"
	"fmt"
	cmd "github.com/BinacsLee/server/cmd/commands"
	"os"
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
	fmt.Println("end")
}
