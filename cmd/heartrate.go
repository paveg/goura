package cmd

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/paveg/goura/api"
	"github.com/paveg/goura/oura"
	"github.com/spf13/cobra"
)

var reqDateTime = &struct {
	startDateTime string
	endDateTime   string
}{}

func heartrateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "heartrate",
		Aliases: []string{"hr"},
		Short:   "Fetch heart rate data",
		Long:    "Fetch heart rate time-series data from the Oura API v2.",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			client, err := api.NewClient(apiBaseURL, &http.Client{}, userAgent, Config.AccessToken)
			if err != nil {
				return err
			}

			period := oura.DateTimePeriod{
				StartDateTime: reqDateTime.startDateTime,
				EndDateTime:   reqDateTime.endDateTime,
			}
			heartrate, err := client.GetHeartRate(ctx, period)
			if err != nil {
				log.Fatalf("fail: %+v", err)
			}
			out(heartrate)

			return nil
		},
	}

	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	cmd.Flags().StringVar(&reqDateTime.startDateTime, "start-datetime", yesterday.Format(time.RFC3339), "start datetime (ISO 8601)")
	cmd.Flags().StringVar(&reqDateTime.endDateTime, "end-datetime", now.Format(time.RFC3339), "end datetime (ISO 8601)")

	return cmd
}
