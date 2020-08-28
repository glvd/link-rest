package linkrest

import (
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestNew(t *testing.T) {
	type args struct {
		port int
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		{
			name: "",
			args: args{
				port: 80,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.port)
			if err != nil {
				t.Fatal(err)
			}
			go got.Start()
			time.Sleep(30 * time.Second)
			err = got.Stop()
			t.Log(err)
		})
	}
}
