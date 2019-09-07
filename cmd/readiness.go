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

func readinessCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "readiness",
		Short: "Fetch readinesses",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			client, err := api.NewClient(apiBaseURL, &http.Client{}, "")
			if err != nil {
				return err
			}
			startDate, endDate, err := initDate()
			if err != nil {
				return err
			}

			datePeriod := api.DatePeriod{StartDate: startDate, EndDate: endDate}
			readiness, err := client.Readiness(ctx, Config.AccessToken, datePeriod)
			if err != nil {
				return err
			}

			var buf bytes.Buffer
			b, _ := json.Marshal(readiness)
			buf.Write(b)
			fmt.Println(buf.String())

			return nil
		},
	}

	return cmd
}
