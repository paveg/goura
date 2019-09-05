package cmd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/paveg/goura/api"
	"github.com/spf13/cobra"
)

func userInfoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "userinfo",
		Short: "get user information",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			client, err := api.NewClient(apiBaseURL, &http.Client{}, "")
			if err != nil {
				return err
			}

			resp, err := client.UserInfo(ctx)
			if err != nil {
				return err
			}
			fmt.Printf("USERINFO: %#v\n", resp)

			return nil
		},
	}
	return cmd
}
