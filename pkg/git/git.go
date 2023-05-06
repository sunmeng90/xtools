package git

import (
	"context"
	"errors"
	"github.com/go-git/go-git/v5"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

func FetchAllWithContext(ctx context.Context, path string) {
	entryChan := make(chan string, 1)
	go scanRepos(ctx, path, entryChan)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work(ctx, entryChan)
		}()
	}
	wg.Wait()
	log.Printf("fetch all finished")
}

func work(ctx context.Context, entryChan <-chan string) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case path, ok := <-entryChan:
			if !ok {
				log.Printf("entry channel closed, work exits")
				return nil
			}
			err := fetchRepo(ctx, path)
			if err != nil {
				log.Printf("failed to fetch repo in %s, reason: %s", path, err)
			}
		}
	}
}

func fetchRepo(ctx context.Context, path string) error {
	r, err := git.PlainOpen(path)
	if err != nil {
		return err
	}
	err = r.FetchContext(ctx, &git.FetchOptions{})
	if err == git.NoErrAlreadyUpToDate {
		return nil
	}
	return err
}

func scanRepos(ctx context.Context, basePath string, entryChan chan<- string) error {
	lstat, err := os.Lstat(basePath)
	if err != nil && !os.IsExist(err) {
		return err
	}
	if !lstat.IsDir() {
		return errors.New("not a folder")
	}
	dir, err := os.ReadDir(basePath)
	if err != nil {
		return err
	}
	for _, entry := range dir {
		select {
		case <-ctx.Done():
			close(entryChan)
			return ctx.Err()
		default:
			if entry.IsDir() {
				entryPath := filepath.Join(basePath, entry.Name())
				entryDir, _ := os.ReadDir(entryPath)
			inner:
				for _, subEntry := range entryDir {
					if subEntry.Name() == ".git" {
						log.Printf("found git repo in %s", entryPath)
						entryChan <- entryPath
						break inner
					}
				}
			}
		}
	}
	log.Printf("finish scan all sub folders. close channel")
	close(entryChan)
	return nil
}
