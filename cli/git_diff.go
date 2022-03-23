package cli

import (
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
)

func GetChangedFiles(directory string) ([]string, error) {
	var err error
	if !filepath.IsAbs(directory) {
		directory, err = filepath.Abs(directory)
		if err != nil {
			return nil, err
		}
	}
	r, err := git.PlainOpen(directory)
	if err != nil {
		return nil, err
	}
	w, err := r.Worktree()
	if err != nil {
		return nil, err
	}
	s, err := w.Status()
	if err != nil {
		return nil, err
	}
	var res []string
	for k, _ := range s {
		if !strings.HasSuffix(k, ".go") {
			continue
		}
		res = append(res, filepath.Join(directory, k))
	}
	return res, nil
}
