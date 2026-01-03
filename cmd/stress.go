package cmd

import (
	"context"
	"log"
	"net/http"

	"github.com/paveg/goura/api"
	"github.com/paveg/goura/oura"
	"github.com/spf13/cobra"
)

func stressCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stress",
		Short: "Fetch daily stress data",
		Long:  "Fetch daily stress and recovery time data from the Oura API v2.",
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
			stress, err := client.GetDailyStress(ctx, datePeriod)
			if err != nil {
				log.Fatalf("fail: %+v", err)
			}
			out(stress)

			return nil
		},
	}

	return cmd
}

func resilienceCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resilience",
		Short: "Fetch daily resilience data",
		Long:  "Fetch daily resilience scores and contributors from the Oura API v2.",
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
			resilience, err := client.GetDailyResilience(ctx, datePeriod)
			if err != nil {
				log.Fatalf("fail: %+v", err)
			}
			out(resilience)

			return nil
		},
	}

	return cmd
}
