package linkrest

import (
	"testing"
	"time"
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
			got := New(tt.args.port)
			go got.Start()
			time.Sleep(5 * time.Second)
			err := got.Stop()
			t.Log(err)
		})
	}
}