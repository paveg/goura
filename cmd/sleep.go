package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/paveg/goura/api"
	"github.com/spf13/cobra"
)

func sleepCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sleeps",
		Short: "fetch sleep",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			client, err := api.NewClient(apiBaseURL, &http.Client{}, "")
			if err != nil {
				return err
			}
			format := "2006-01-02"
			start, err := time.Parse(format, "2019-08-01")
			if err != nil {
				return err
			}
			end, err := time.Parse(format, "2019-09-01")
			if err != nil {
				return err
			}
			dp := api.DatePeriod{StartDate: start, EndDate: end}
			resp, err := client.Sleep(ctx, Config.AccessToken, dp)
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
