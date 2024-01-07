package ltlsp

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lt-lsp)")
	viper.BindPFlag("uername", RootCmd.PersistentFlags().Lookup("username"))
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".lt-lsp")
		viper.SetConfigType("yaml")
	}

	viper.ReadInConfig()
}

var RootCmd = &cobra.Command{
	Use:   "lt-lsp",
	Short: "lt-lsp - a simple LSP to wrap LanguageTool",
	Long:  `A simple Language Server to wrap LanguageTool`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
