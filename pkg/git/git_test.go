package git

import (
	"context"
	"testing"
	"time"
)

func TestFetchAllWithContext(t *testing.T) {
	type args struct {
		path    string
		timeout string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "current folder",
			args: args{
				path:    "I:\\ws\\GitHub\\sunmeng90\\leetcode-go",
				timeout: "5s",
			},
		},
		{
			name: "empty base",
			args: args{
				path:    "I:\\ws\\GitHub\\sunmeng90\\go\\xtools\\pkg\\git",
				timeout: "5s",
			},
		},
		{
			name: "github base",
			args: args{
				path:    "I:\\ws\\GitHub\\",
				timeout: "15s",
			},
		},
		{
			name: "not exist folder",
			args: args{
				path:    "asfdsdafsd",
				timeout: "5s",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeoutDuration, _ := time.ParseDuration(tt.args.timeout)
			ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
			defer cancel()
			FetchAllWithContext(ctx, tt.args.path)
		})
	}
}
