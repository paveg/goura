package api_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/paveg/goura/api"
)

func TestClient_UserInfo(t *testing.T) {
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
	mux, mockServerURL := api.NewMockServer()
	client := api.NewTestClient(mockServerURL)

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
