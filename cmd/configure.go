package cmd

import (
	"github.com/paveg/goura/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func configCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "configure",
		Short: "Fetch your access_token",
		RunE: func(cmd *cobra.Command, args []string) error {
			token, err := api.FetchAccessToken(Config.ClientID, Config.ClientSecret)
			if err != nil {
				return err
			}

			viper.Set("AccessToken", token)
			if err := viper.WriteConfig(); err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}

func init() {}
