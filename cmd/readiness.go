package cmd

import (
	"context"
	"log"
	"net/http"

	"github.com/paveg/goura/api"
	"github.com/paveg/goura/oura"
	"github.com/spf13/cobra"
)

func readinessCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "readiness",
		Short: "Fetch readinesses",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			client, err := api.NewClient(apiBaseURL, &http.Client{}, "", Config.AccessToken)
			if err != nil {
				return err
			}
			startDate, endDate, err := initDate()
			if err != nil {
				return err
			}

			datePeriod := oura.DatePeriod{StartDate: startDate, EndDate: endDate}
			readiness, err := client.GetReadiness(ctx, datePeriod)
			if err != nil {
				log.Fatalf("fail: %+v", err)
			}
			out(readiness)

			return nil
		},
	}

	return cmd
}
