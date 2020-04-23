package main

import (
	"github.com/spf13/cobra"
	"log"
)

func run() {
	var (
		bootstrapperCmd = &cobra.Command{
			Use: "go-bootstrapper",
			Version: "0.1.0",
			Args: cobra.NoArgs,
			Short: "A Go project bootstrapper",
			Long: "",
			Example: "",
			Run: func(ccmd *cobra.Command, args []string){},
		}

		newCmd = &cobra.Command{
			Use: "new",
			Short: "Bootstrap a new project",
			Long: "",
			Args: cobra.NoArgs,
			Example: "",
			Run: func(ccmd *cobra.Command, args []string){},
		}

		cliCmd = &cobra.Command{
			Use: "cli",
			Short: "Bootstrap a new CLI project",
			Long: "",
			Args: cobra.NoArgs,
			Example: "",
			Run: func(ccmd *cobra.Command, args []string){},
		}
	)

	bootstrapperCmd.AddCommand(newCmd)
	newCmd.AddCommand(cliCmd)

	if err := bootstrapperCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	run()
}