package git

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"sync"
)

type Repo struct {
	LocalPath string
}

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
				if entry.Name() == ".git" {
					log.Infof("found git repo in %s", basePath)
					entryChan <- basePath
				}
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

// CheckoutBranchOrHash create a new branch based on another branch or revision
func CheckoutBranchOrHash(repoPath string, branchOrHash string) error {
	r, err := openRepoInPath(repoPath)
	if err != nil {
		log.Errorf("failed to open repo in %s, err: %s", repoPath, err)
		return err
	}
	return CheckoutBranchWithRepo(r, branchOrHash)
}

func openRepoInPath(repoPath string) (*git.Repository, error) {
	path, err := resolvePath(repoPath)
	if err != nil {
		return nil, err
	}
	r, err := git.PlainOpenWithOptions(path, &git.PlainOpenOptions{
		DetectDotGit:          false,
		EnableDotGitCommonDir: false,
	})
	return r, err
}

func resolvePath(repoPath string) (string, error) {
	if repoPath == "" {
		repoPath = "."
	}
	path, err := filepath.Abs(repoPath)
	if err != nil {
		return "", err
	}
	return path, nil
}

func CreateBranch(repoPath, targetBranchName, srcRev string) error {
	r, err := openRepoInPath(repoPath)
	if err != nil {
		log.Errorf("failed to open reop in %s, err: %s", repoPath, err)
		return err
	}

	hash, err := r.ResolveRevision(plumbing.Revision(srcRev))
	if err != nil {
		log.Errorf("failed to resolve revision %s. err: %s", srcRev, err)
		return err
	}

	refForBranch := plumbing.NewHashReference(
		plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", targetBranchName)), *hash)
	err = r.Storer.SetReference(refForBranch)
	if err == nil {
		log.Infof("create rev " + refForBranch.String())
	}
	return err
}

func CreateBranchForRepo(repo Repo, srcRev, targetBranchName string) error {
	return CreateBranch(repo.LocalPath, targetBranchName, srcRev)
}

func CheckoutBranchOrHashForRepo(repo Repo, branchOrHash string) error {
	r, err := git.PlainOpenWithOptions(repo.LocalPath, &git.PlainOpenOptions{
		DetectDotGit:          false,
		EnableDotGitCommonDir: false,
	})
	if err != nil {
		log.Errorf("failed to open reop in %s, err: %s", repo, err)
		return err
	}
	return CheckoutBranchWithRepo(r, branchOrHash)
}

func CheckoutBranchWithRepo(r *git.Repository, branchOrHash string) error {
	w, err := r.Worktree()
	if err != nil {
		return err
	}

	branchRef := fmt.Sprintf("refs/heads/%s", branchOrHash)
	_, err = r.Storer.Reference(plumbing.ReferenceName(branchRef))
	if err == nil {
		log.Infof("checkout rev %s", branchRef)
		return w.Checkout(&git.CheckoutOptions{
			Branch: plumbing.ReferenceName(branchRef),
			Create: false,
			Keep:   true,
		})
	}

	hash, err := r.ResolveRevision(plumbing.Revision(branchOrHash))
	if err != nil {
		log.Errorf("failed to resolve revision [%s]. err: %s", branchOrHash, err)
		return err
	}
	return w.Checkout(&git.CheckoutOptions{Hash: *hash, Create: false})
}
