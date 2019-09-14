package cmd

import (
	"context"
	"log"
	"net/http"

	"github.com/paveg/goura/api"
	"github.com/paveg/goura/oura"
	"github.com/spf13/cobra"
)

func sleepCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sleep",
		Short: "Fetch sleep",
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
			sleeps, err := client.GetSleep(ctx, datePeriod)
			if err != nil {
				log.Fatalf("fail: %+v", err)
			}
			out(sleeps)

			return nil
		},
	}

	return cmd
}
