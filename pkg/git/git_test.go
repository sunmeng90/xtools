package git

import (
	"context"
	"path/filepath"
	"testing"
	"time"
)

const (
	testGitBase = "C:\\tmp\\github\\"
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
				path:    filepath.Join(testGitBase, "\\repo1"),
				timeout: "5s",
			},
		},
		{
			name: "empty base",
			args: args{
				path:    filepath.Join(testGitBase, "\\repo1"),
				timeout: "5s",
			},
		},
		{
			name: "github base",
			args: args{
				path:    testGitBase,
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

func TestCreateBranch(t *testing.T) {
	type args struct {
		repo       Repo
		rev        string
		branchName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "create b1",
			args: args{
				repo:       Repo{LocalPath: filepath.Join(testGitBase, "\\repo1")},
				rev:        "b1~2",
				branchName: "b3",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateBranchForRepo(tt.args.repo, tt.args.rev, tt.args.branchName); (err != nil) != tt.wantErr {
				t.Errorf("CreateBranchForRepo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckoutBranch(t *testing.T) {
	type args struct {
		repo Repo
		rev  string
	}
	sampleGitRepoPath := filepath.Join(testGitBase, "\\repo1")
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "checkout b1",
			args: args{
				repo: Repo{LocalPath: sampleGitRepoPath},
				rev:  "b1",
			},
			wantErr: false,
		},
		{
			name: "checkout 6d9bac58001661a9ceeba4b3f6d4315280bc50ce",
			args: args{
				repo: Repo{LocalPath: sampleGitRepoPath},
				rev:  "6d9bac58001661a9ceeba4b3f6d4315280bc50ce",
			},
			wantErr: false,
		},
		{
			name: "checkout main^1",
			args: args{
				repo: Repo{LocalPath: sampleGitRepoPath},
				rev:  "main^1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckoutBranchOrHashForRepo(tt.args.repo, tt.args.rev); (err != nil) != tt.wantErr {
				t.Errorf("CheckoutBranchOrHashForRepo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
