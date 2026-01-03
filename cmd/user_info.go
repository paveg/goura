package cmd

import (
	"context"
	"log"
	"net/http"

	"github.com/paveg/goura/api"
	"github.com/spf13/cobra"
)

func userInfoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "userinfo",
		Aliases: []string{"personal-info", "me"},
		Short:   "Fetch user personal information",
		Long:    "Fetch user personal information from the Oura API v2 (age, weight, height, biological sex, email).",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			client, err := api.NewClient(apiBaseURL, &http.Client{}, userAgent, Config.AccessToken)
			if err != nil {
				return err
			}

			userinfo, err := client.GetPersonalInfo(ctx)
			if err != nil {
				log.Fatalf("fail: %+v", err)
			}
			out(userinfo)

			return nil
		},
	}
	return cmd
}
