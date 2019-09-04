package api_test

import (
	"context"
	"fmt"
	"github.com/paveg/goura/api"
	"net/http"
	"reflect"
	"testing"
)

func TestClient_UserInfo(t *testing.T) {
	tests := []struct {
		res  string
		want *api.UserInfoResponse
	}{
		{
			res: `{
  "age": 27,
  "weight": 80,
  "gender": "male",
  "email": "john.doe@the.domain"
}`,
			want: &api.UserInfoResponse{
				Age:    27,
				Weight: 80,
				Gender: "male",
				Email:  "john.doe@the.domain",
			},
		},
	}
	mux, mockServerURL := api.NewMockServer()
	client := api.NewTestClient(mockServerURL)

	for _, tt := range tests {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprintf(w, tt.res)
		})

		got, err := client.UserInfo(context.Background())

		if err != nil {
			t.Fatalf("UserInfo was failed: got = %+v, err = %+v", got, err)
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got : %#v, want: %#v", got, tt.want)
		}
	}
}
