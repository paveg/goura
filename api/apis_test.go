package api_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/paveg/goura/api"
	"github.com/paveg/goura/oura"
)

func TestClient_GetPersonalInfo(t *testing.T) {
	mux, client := initTest(t)
	tests := []struct {
		res  string
		want *oura.PersonalInfo
	}{
		{
			res: `{
  "id": "abc123",
  "age": 27,
  "weight": 80.2,
  "height": 180,
  "biological_sex": "male",
  "email": "john.doe@the.domain"
}`,
			want: &oura.PersonalInfo{
				ID:            "abc123",
				Age:           27,
				Weight:        80.2,
				Height:        180.0,
				BiologicalSex: "male",
				Email:         "john.doe@the.domain",
			},
		},
	}

	for _, tt := range tests {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprintf(w, tt.res)
		})

		got, err := client.GetPersonalInfo(context.Background())

		if err != nil {
			t.Fatalf("GetPersonalInfo was failed: got = %+v, err = %+v", got, err)
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got : %+v, want: %+v", got, tt.want)
		}
	}
}

func TestClient_GetDailyActivity(t *testing.T) {
	mux, client := initTest(t)
	score := 85
	tests := []struct {
		datePeriod oura.DatePeriod
		res        string
		want       *oura.PaginatedResponse[oura.DailyActivity]
	}{
		{
			datePeriod: oura.DatePeriod{StartDate: "2024-01-01", EndDate: "2024-01-01"},
			res: `{
  "data": [
    {
      "id": "activity-123",
      "day": "2024-01-01",
      "score": 85,
      "steps": 10000,
      "active_calories": 500,
      "total_calories": 2000
    }
  ],
  "next_token": null
}`,
			want: &oura.PaginatedResponse[oura.DailyActivity]{
				Data: []oura.DailyActivity{
					{
						ID:            "activity-123",
						Day:           "2024-01-01",
						Score:         &score,
						Steps:         10000,
						ActiveCalories: 500,
						TotalCalories: 2000,
					},
				},
				NextToken: nil,
			},
		},
	}

	for _, tt := range tests {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprintf(w, tt.res)
		})

		got, err := client.GetDailyActivity(context.Background(), tt.datePeriod)

		if err != nil {
			t.Fatalf("GetDailyActivity was failed: got = %+v, err = %+v", got, err)
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got : %+v, want: %+v", got, tt.want)
		}
	}
}

func TestClient_GetDailyReadiness(t *testing.T) {
	mux, client := initTest(t)
	score := 90
	tests := []struct {
		datePeriod oura.DatePeriod
		res        string
		want       *oura.PaginatedResponse[oura.DailyReadiness]
	}{
		{
			datePeriod: oura.DatePeriod{StartDate: "2024-01-01", EndDate: "2024-01-01"},
			res: `{
  "data": [
    {
      "id": "readiness-123",
      "day": "2024-01-01",
      "score": 90,
      "timestamp": "2024-01-01T07:00:00+00:00"
    }
  ],
  "next_token": null
}`,
			want: &oura.PaginatedResponse[oura.DailyReadiness]{
				Data: []oura.DailyReadiness{
					{
						ID:        "readiness-123",
						Day:       "2024-01-01",
						Score:     &score,
						Timestamp: "2024-01-01T07:00:00+00:00",
					},
				},
				NextToken: nil,
			},
		},
	}

	for _, tt := range tests {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprintf(w, tt.res)
		})

		got, err := client.GetDailyReadiness(context.Background(), tt.datePeriod)

		if err != nil {
			t.Fatalf("GetDailyReadiness was failed: got = %+v, err = %+v", got, err)
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got : %+v, want: %+v", got, tt.want)
		}
	}
}

func TestClient_GetDailySleep(t *testing.T) {
	mux, client := initTest(t)
	score := 88
	tests := []struct {
		datePeriod oura.DatePeriod
		res        string
		want       *oura.PaginatedResponse[oura.DailySleep]
	}{
		{
			datePeriod: oura.DatePeriod{StartDate: "2024-01-01", EndDate: "2024-01-01"},
			res: `{
  "data": [
    {
      "id": "sleep-123",
      "day": "2024-01-01",
      "score": 88,
      "timestamp": "2024-01-01T07:30:00+00:00"
    }
  ],
  "next_token": null
}`,
			want: &oura.PaginatedResponse[oura.DailySleep]{
				Data: []oura.DailySleep{
					{
						ID:        "sleep-123",
						Day:       "2024-01-01",
						Score:     &score,
						Timestamp: "2024-01-01T07:30:00+00:00",
					},
				},
				NextToken: nil,
			},
		},
	}

	for _, tt := range tests {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprintf(w, tt.res)
		})

		got, err := client.GetDailySleep(context.Background(), tt.datePeriod)

		if err != nil {
			t.Fatalf("GetDailySleep was failed: got = %+v, err = %+v", got, err)
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got : %+v, want: %+v", got, tt.want)
		}
	}
}

func TestClient_GetSleep(t *testing.T) {
	mux, client := initTest(t)
	avgHR := 55.5
	avgBreath := 14.2
	tests := []struct {
		datePeriod oura.DatePeriod
		res        string
		want       *oura.PaginatedResponse[oura.Sleep]
	}{
		{
			datePeriod: oura.DatePeriod{StartDate: "2024-01-01", EndDate: "2024-01-01"},
			res: `{
  "data": [
    {
      "id": "sleep-period-123",
      "day": "2024-01-01",
      "bedtime_start": "2024-01-01T23:00:00+00:00",
      "bedtime_end": "2024-01-02T07:00:00+00:00",
      "average_heart_rate": 55.5,
      "average_breath": 14.2,
      "type": "long_sleep"
    }
  ],
  "next_token": null
}`,
			want: &oura.PaginatedResponse[oura.Sleep]{
				Data: []oura.Sleep{
					{
						ID:               "sleep-period-123",
						Day:              "2024-01-01",
						BedtimeStart:     "2024-01-01T23:00:00+00:00",
						BedtimeEnd:       "2024-01-02T07:00:00+00:00",
						AverageHeartRate: &avgHR,
						AverageBreath:    &avgBreath,
						Type:             "long_sleep",
					},
				},
				NextToken: nil,
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

func TestClient_GetHeartRate(t *testing.T) {
	mux, client := initTest(t)
	tests := []struct {
		period oura.DateTimePeriod
		res    string
		want   *oura.PaginatedResponse[oura.HeartRate]
	}{
		{
			period: oura.DateTimePeriod{StartDateTime: "2024-01-01T00:00:00Z", EndDateTime: "2024-01-01T23:59:59Z"},
			res: `{
  "data": [
    {
      "bpm": 72,
      "source": "awake",
      "timestamp": "2024-01-01T10:30:00+00:00"
    },
    {
      "bpm": 65,
      "source": "rest",
      "timestamp": "2024-01-01T03:00:00+00:00"
    }
  ],
  "next_token": null
}`,
			want: &oura.PaginatedResponse[oura.HeartRate]{
				Data: []oura.HeartRate{
					{
						BPM:       72,
						Source:    "awake",
						Timestamp: "2024-01-01T10:30:00+00:00",
					},
					{
						BPM:       65,
						Source:    "rest",
						Timestamp: "2024-01-01T03:00:00+00:00",
					},
				},
				NextToken: nil,
			},
		},
	}

	for _, tt := range tests {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprintf(w, tt.res)
		})

		got, err := client.GetHeartRate(context.Background(), tt.period)

		if err != nil {
			t.Fatalf("GetHeartRate was failed: got = %+v, err = %+v", got, err)
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got : %+v, want: %+v", got, tt.want)
		}
	}
}

func TestClient_GetWorkouts(t *testing.T) {
	mux, client := initTest(t)
	calories := 350.5
	distance := 5000.0
	tests := []struct {
		datePeriod oura.DatePeriod
		res        string
		want       *oura.PaginatedResponse[oura.Workout]
	}{
		{
			datePeriod: oura.DatePeriod{StartDate: "2024-01-01", EndDate: "2024-01-01"},
			res: `{
  "data": [
    {
      "id": "workout-123",
      "activity": "running",
      "calories": 350.5,
      "day": "2024-01-01",
      "distance": 5000.0,
      "intensity": "moderate",
      "source": "manual"
    }
  ],
  "next_token": null
}`,
			want: &oura.PaginatedResponse[oura.Workout]{
				Data: []oura.Workout{
					{
						ID:        "workout-123",
						Activity:  "running",
						Calories:  &calories,
						Day:       "2024-01-01",
						Distance:  &distance,
						Intensity: "moderate",
						Source:    "manual",
					},
				},
				NextToken: nil,
			},
		},
	}

	for _, tt := range tests {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprintf(w, tt.res)
		})

		got, err := client.GetWorkouts(context.Background(), tt.datePeriod)

		if err != nil {
			t.Fatalf("GetWorkouts was failed: got = %+v, err = %+v", got, err)
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got : %+v, want: %+v", got, tt.want)
		}
	}
}

func TestClient_GetDailyStress(t *testing.T) {
	mux, client := initTest(t)
	stressHigh := 120
	recoveryHigh := 300
	tests := []struct {
		datePeriod oura.DatePeriod
		res        string
		want       *oura.PaginatedResponse[oura.DailyStress]
	}{
		{
			datePeriod: oura.DatePeriod{StartDate: "2024-01-01", EndDate: "2024-01-01"},
			res: `{
  "data": [
    {
      "id": "stress-123",
      "day": "2024-01-01",
      "stress_high": 120,
      "recovery_high": 300,
      "day_summary": "normal"
    }
  ],
  "next_token": null
}`,
			want: &oura.PaginatedResponse[oura.DailyStress]{
				Data: []oura.DailyStress{
					{
						ID:           "stress-123",
						Day:          "2024-01-01",
						StressHigh:   &stressHigh,
						RecoveryHigh: &recoveryHigh,
						DaySummary:   "normal",
					},
				},
				NextToken: nil,
			},
		},
	}

	for _, tt := range tests {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprintf(w, tt.res)
		})

		got, err := client.GetDailyStress(context.Background(), tt.datePeriod)

		if err != nil {
			t.Fatalf("GetDailyStress was failed: got = %+v, err = %+v", got, err)
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
