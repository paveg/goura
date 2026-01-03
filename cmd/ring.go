package cmd

import (
	"context"
	"log"
	"net/http"

	"github.com/paveg/goura/api"
	"github.com/paveg/goura/oura"
	"github.com/spf13/cobra"
)

func ringCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ring",
		Aliases: []string{"ring-config"},
		Short:   "Fetch ring configuration",
		Long:    "Fetch ring hardware and firmware details from the Oura API v2.",
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
			ring, err := client.GetRingConfiguration(ctx, datePeriod)
			if err != nil {
				log.Fatalf("fail: %+v", err)
			}
			out(ring)

			return nil
		},
	}

	return cmd
}
