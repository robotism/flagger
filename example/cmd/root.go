package cmd

import (
	"log"
	"os"

	"github.com/robotism/flagger"
	"github.com/robotism/flagger/example/config"
	"github.com/spf13/cobra"
)

var (
	f = flagger.New()
	c = &config.AppConfig{}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "example",
	Short: "a flagger example",
	Long:  `a flagger example`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("%+v\n", c)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	f.UseFlags(rootCmd.Flags())
	f.UseConfigFileArgDefault()
	err := f.Parse(c)
	if err != nil {
		panic(err)
	}
}
