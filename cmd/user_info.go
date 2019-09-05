package cmd

import (
	"bytes"
	"context"
	"encoding/json"
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

			resp, err := client.UserInfo(ctx, Config.AccessToken)
			if err != nil {
				return err
			}
			var buf bytes.Buffer
			b, _ := json.Marshal(resp)
			buf.Write(b)
			fmt.Println(buf.String())

			return nil
		},
	}
	return cmd
}
