package cmd

import (
	"context"
	"log"
	"net/http"

	"github.com/paveg/goura/api"
	"github.com/paveg/goura/oura"
	"github.com/spf13/cobra"
)

func workoutCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "workout",
		Aliases: []string{"workouts"},
		Short:   "Fetch workout data",
		Long:    "Fetch workout data including activity type, calories, distance, and intensity from the Oura API v2.",
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
			workouts, err := client.GetWorkouts(ctx, datePeriod)
			if err != nil {
				log.Fatalf("fail: %+v", err)
			}
			out(workouts)

			return nil
		},
	}

	return cmd
}
