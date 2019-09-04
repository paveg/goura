package command_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/mattn/go-shellwords"
	"github.com/paveg/goura/command"
)

func TestExecute(t *testing.T) {
	tests := []struct {
		command string
		want    string
		wantErr string
	}{
		{
			command: "goura version",
			want:    "",
			wantErr: "",
		},
	}

	for _, tt := range tests {
		buf := new(bytes.Buffer)
		cmd := command.NewCommandRoot()
		cmd.SetOut(buf)
		cmdArgs, err := shellwords.Parse(tt.command)
		if err != nil {
			t.Fatalf("args parse error: %+v\n", err)
		}

		cmd.SetArgs(cmdArgs[1:])
		if err := cmd.Execute(); err != nil {
			got := err.Error()
			if tt.want != got {
				t.Errorf("unexpected error response - want: %+v, got: %+v\n", tt.want, got)
			}
		}

		got := buf.String()
		fmt.Println("\n=== Buffer End")
		if tt.want != got {
			t.Errorf("unexpected response - want: %+v, got: %+v\n", tt.want, got)
		}
	}
}
