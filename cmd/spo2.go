package cmd

import (
	"context"
	"log"
	"net/http"

	"github.com/paveg/goura/api"
	"github.com/paveg/goura/oura"
	"github.com/spf13/cobra"
)

func spo2Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "spo2",
		Aliases: []string{"oxygen"},
		Short:   "Fetch daily SpO2 (blood oxygen) data",
		Long:    "Fetch daily SpO2 (blood oxygen saturation) data from the Oura API v2. Requires Gen 3 ring or later.",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			client, err := api.NewClient(apiBaseURL, &http.Client{}, userAgent, Config.AccessToken)
			if err != nil {
				return err
			}
			startDate, endDate, err := initDate()
			if err != nil {
				return err
			}

			datePeriod := oura.DatePeriod{StartDate: startDate, EndDate: endDate}
			spo2, err := client.GetDailySpo2(ctx, datePeriod)
			if err != nil {
				log.Fatalf("fail: %+v", err)
			}
			out(spo2)

			return nil
		},
	}

	return cmd
}
