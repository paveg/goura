package api_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/paveg/goura/api"
)

func TestClient_GetUserInfo(t *testing.T) {
	mux, client := initTest(t)
	tests := []struct {
		res  string
		want *api.UserInfo
	}{
		{
			res: `{
  "age": 27,
  "weight": 80.2,
  "height": 180,
  "gender": "male",
  "email": "john.doe@the.domain",
  "user_id": "abc"
}`,
			want: &api.UserInfo{
				Age:    27,
				Weight: 80.2,
				Height: 180.0,
				Gender: "male",
				Email:  "john.doe@the.domain",
				UserID: "abc",
			},
		},
	}

	for _, tt := range tests {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprintf(w, tt.res)
		})

		got, err := client.GetUserInfo(context.Background())

		if err != nil {
			t.Fatalf("GetUserInfo was failed: got = %+v, err = %+v", got, err)
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got : %+v, want: %+v", got, tt.want)
		}
	}
}

func TestClient_GetActivity(t *testing.T) {
	mux, client := initTest(t)
	tests := []struct {
		datePeriod api.DatePeriod
		res        string
		want       *api.Activities
	}{
		{
			datePeriod: api.DatePeriod{StartDate: "2015-01-01", EndDate: "2015-01-01"},
			res: `{ "activity": 
	[
		{
			"summary_date": "2015-01-01",
			"average_met": 1.0
		}
	]
}`,
			want: &api.Activities{
				Activity: []api.Activity{
					{
						SummaryDate: "2015-01-01",
						AverageMet:    1.0,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprintf(w, tt.res)
		})

		got, err := client.GetActivity(context.Background(), tt.datePeriod)

		if err != nil {
			t.Fatalf("GetActivity was failed: got = %+v, err = %+v", got, err)
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got : %+v, want: %+v", got, tt.want)
		}
	}
}

func TestClient_GetReadiness(t *testing.T) {
	mux, client := initTest(t)
	tests := []struct {
		datePeriod api.DatePeriod
		res        string
		want       *api.Readinesses
	}{
		{
			datePeriod: api.DatePeriod{StartDate: "2015-01-01", EndDate: "2015-01-01"},
			res: `{ "readiness": 
	[
		{
			"summary_date": "2015-01-01",
			"period_id": 1
		}
	]
}`,
			want: &api.Readinesses{
				Readiness: []api.Readiness{
					{
						SummaryDate: "2015-01-01",
						PeriodID:    1,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprintf(w, tt.res)
		})

		got, err := client.GetReadiness(context.Background(), tt.datePeriod)

		if err != nil {
			t.Fatalf("GetReadiness was failed: got = %+v, err = %+v", got, err)
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got : %+v, want: %+v", got, tt.want)
		}
	}
}

func TestClient_GetSleep(t *testing.T) {
	mux, client := initTest(t)
	tests := []struct {
		datePeriod api.DatePeriod
		res        string
		want       *api.SleepPeriods
	}{
		{
			datePeriod: api.DatePeriod{StartDate: "2015-01-01", EndDate: "2015-01-01"},
			res: `{ "sleep": 
	[
		{
			"summary_date": "2015-01-01",
			"period_id": 1
		}
	]
}`,
			want: &api.SleepPeriods{
				Sleeps: []api.Sleep{
					{
						SummaryDate: "2015-01-01",
						PeriodID:    1,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprintf(w, tt.res)
		})

		got, err := client.GetSleep(context.Background(), tt.datePeriod)

		if err != nil {
			t.Fatalf("GetSleep was failed: got = %+v, err = %+v", got, err)
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got : %+v, want: %+v", got, tt.want)
		}
	}
}

func initTest(t *testing.T) (*http.ServeMux, *api.Client) {
	t.Helper()
	mux, mockServerURL := api.NewMockServer()
	client := api.NewTestClient(mockServerURL)

	return mux, client
}
