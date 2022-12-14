package main

import (
	"github.com/fajarardiyanto/boilerplate-rest-api/cmd/api"
	"github.com/fajarardiyanto/boilerplate-rest-api/cmd/generate"
	"github.com/fajarardiyanto/boilerplate-rest-api/config"
	"github.com/fajarardiyanto/boilerplate-rest-api/internal/dto"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	if err := Run(os.Args[1:]); err != nil {
		config.GetLogger().Error("Unable to run the command %s ", err.Error())
	}
}

var rootCmd = &cobra.Command{
	Use:   dto.GetConfig().Name,
	Short: dto.GetConfig().Name,
}

func init() {
	rootCmd.AddCommand(api.CmdAPI)
	rootCmd.AddCommand(generate.CmdConfig)
}

// Run function lets you run the commands
func Run(args []string) error {
	rootCmd.SetArgs(args)
	return rootCmd.Execute()
}
