package git

import (
	"context"
	"errors"
	"github.com/go-git/go-git/v5"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"sync"
)

func FetchAllWithContext(ctx context.Context, path string) {
	entryChan := make(chan string, 1)
	go scanRepos(ctx, path, entryChan)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			work(ctx, entryChan, idx)
			log.Debugf("worker %d is done", idx)
		}(i)
	}
	wg.Wait()
	log.Info("fetch all finished")
}

func work(ctx context.Context, entryChan <-chan string, idx int) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case path, ok := <-entryChan:
			if !ok {
				log.Debugf("repository folder channel closed, work %d exits", idx)
				return nil
			}
			err := fetchRepo(ctx, path)
			if err != nil {
				log.Errorf("failed to fetch repo in %s, reason: %s", path, err)
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
		log.Debugf("repo in %s is update to date", path)
		return nil
	}
	return err
}

func scanRepos(ctx context.Context, basePath string, entryChan chan<- string) error {
	defer close(entryChan)
	lstat, err := os.Lstat(basePath)
	if err != nil && !os.IsExist(err) {
		log.Errorf("%s doesn't exist, error: %s", basePath, err)
		return err
	}
	if !lstat.IsDir() {
		log.Warnf("%s is not a dir", basePath)
		return errors.New("not a folder")
	}
	dir, err := os.ReadDir(basePath)
	if err != nil {
		log.Errorf("failed to read dir %s", basePath)
		return err
	}
	for _, entry := range dir {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if entry.IsDir() {
				entryPath := filepath.Join(basePath, entry.Name())
				entryDir, _ := os.ReadDir(entryPath)
			inner:
				for _, subEntry := range entryDir {
					if subEntry.Name() == ".git" {
						log.Infof("found git repo in %s", entryPath)
						entryChan <- entryPath
						break inner
					}
				}
			}
		}
	}
	log.Info("finish scan all sub folders. close channel")
	return nil
}
