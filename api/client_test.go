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

func TestClient_PersonalInfo(t *testing.T) {
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
	mux, mockServerURL := api.NewMockServer()
	client := api.NewTestClient(mockServerURL)

	for _, tt := range tests {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprint(w, tt.res)
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
