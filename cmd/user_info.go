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
		Use:   "userinfo",
		Short: "Fetch user information",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			client, err := api.NewClient(apiBaseURL, &http.Client{}, "", Config.AccessToken)
			if err != nil {
				return err
			}

			userinfo, err := client.GetUserInfo(ctx)
			if err != nil {
				log.Fatalf("fail: %+v", err)
			}
			out(userinfo)

			return nil
		},
	}
	return cmd
}
