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

type requiredDate struct {
	start  string
	end    string
	target string
}

var reqDate = &requiredDate{}

const dateFormat = "2006-01-02"

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
			startDate, endDate, err := initDate()
			if err != nil {
				return err
			}

			datePeriod := api.DatePeriod{StartDate: startDate, EndDate: endDate}
			resp, err := client.Sleep(ctx, Config.AccessToken, datePeriod)
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

func initDate() (string, string, error) {
	if reqDate.target != "" {
		reqDate.start = reqDate.target
		reqDate.end = reqDate.target
	}

	startDate, err := time.Parse(dateFormat, reqDate.start)
	if err != nil {
		return "", "", err
	}
	endDate, err := time.Parse(dateFormat, reqDate.end)
	if err != nil {
		return "", "", err
	}

	return startDate.Format(dateFormat), endDate.Format(dateFormat), err
}
