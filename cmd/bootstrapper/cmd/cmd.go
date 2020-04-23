package main

import (
	"github.com/j4ng5y/awesomeProject1/pkg/bootstrapper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
)

func run() {
	var (
		vCfg = viper.New()
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
			Use: "new-project",
			Short: "Bootstrap a new project",
			Long: "",
			Example: "",
			Run: func(ccmd *cobra.Command, args []string){
				if len(args) == 1 {
					vCfg.Set("project.name", args[0])
					log.Printf("Bootstrapping the new project: %s", vCfg.GetString("project.name"))
					b := bootstrapper.New(vCfg)
					if errs := b.GenerateDirectories(); errs != nil {
						log.Fatalf("errors occurred: %+v", errs)
					}
					if errs := b.GenerateFiles(); errs != nil {
						log.Fatalf("errors occurred: %+v", errs)
					}
					workingDir, _ := os.Getwd()
					log.Printf("The project has been bootstrapped, you can navigate to it here: %s", path.Join(workingDir, vCfg.GetString("project.name")))
				} else if len(args) >1 {
					for _, v := range args {
						vCfg.Set("project.name", v)
						log.Printf("Bootstrapping the new project: %s", vCfg.GetString("project.name"))
						b := bootstrapper.New(vCfg)
						if errs := b.GenerateDirectories(); errs != nil {
							log.Fatalf("errors occurred: %+v", errs)
						}
						if errs := b.GenerateFiles(); errs != nil {
							log.Fatalf("errors occurred: %+v", errs)
						}
						workingDir, _ := os.Getwd()
						log.Printf("The project has been bootstrapped, you can navigate to it here: %s", path.Join(workingDir, vCfg.GetString("project.name")))
					}
				} else {
					log.Fatal("at least one argument is required")
				}
			},
		}
	)

	bootstrapperCmd.AddCommand(newCmd)

	if err := bootstrapperCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	run()
}