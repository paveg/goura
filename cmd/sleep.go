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
		Short: "Fetch daily sleep scores",
		Long:  "Fetch daily sleep scores from the Oura API v2.",
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
			sleeps, err := client.GetDailySleep(ctx, datePeriod)
			if err != nil {
				log.Fatalf("fail: %+v", err)
			}
			out(sleeps)

			return nil
		},
	}

	return cmd
}

func sleepPeriodsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sleep-periods",
		Short: "Fetch detailed sleep period data",
		Long:  "Fetch detailed sleep period data including heart rate, HRV, and sleep stages from the Oura API v2.",
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

func sleepTimeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sleep-time",
		Short: "Fetch sleep time recommendations",
		Long:  "Fetch optimal bedtime recommendations from the Oura API v2.",
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
			sleepTime, err := client.GetSleepTime(ctx, datePeriod)
			if err != nil {
				log.Fatalf("fail: %+v", err)
			}
			out(sleepTime)

			return nil
		},
	}

	return cmd
}
