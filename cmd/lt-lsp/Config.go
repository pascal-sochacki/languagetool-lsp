package ltlsp

import (
	"context"
	"errors"

	"github.com/pascal-sochacki/languagetool-lsp/pkg/languagetool"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func init() {
	RootCmd.AddCommand(Configure)
}

var Configure = &cobra.Command{
	Use:   "configure <username> <api-token>",
	Short: "Configure the LSP (set api Token and username)",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(2)(cmd, args); err != nil {
			return err
		}
		client := languagetool.NewClient(zap.NewNop())
		client = client.WithCredentials(languagetool.Credentials{
			Username: args[0],
			ApiToken: args[1],
		})
		// make a request to check if the credentials are valid
		result, err := client.CheckText(context.Background(), "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.")

		if err != nil {
			return err
		}

		if !result.Software.Premium {
			return errors.New("Tried to call the api but don't received a premium response")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("username", args[0])
		viper.Set("apiToken", args[1])
		viper.WriteConfig()
	},
}
