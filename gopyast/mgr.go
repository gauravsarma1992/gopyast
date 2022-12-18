package gopyast

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type (
	Manager struct {
		BaseDir     string
		WorkerCount int

		wg *sync.WaitGroup
	}
)

func New() (mgr *Manager, err error) {
	mgr = &Manager{
		BaseDir:     fmt.Sprintf("%s/app", os.Getenv("SEARCH_MAIN_DIR")),
		WorkerCount: 8,
		wg:          &sync.WaitGroup{},
	}
	return
}

func (mgr *Manager) GetFiles() (filePaths []string, err error) {
	if err = filepath.Walk(mgr.BaseDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.Index(path, ".py") == -1 || strings.Index(path, "test") != -1 {
				return err
			}
			filePaths = append(filePaths, path)
			return nil
		}); err != nil {
		return
	}
	return
}

func (mgr *Manager) Run() (err error) {
	var (
		filePaths []string
	)
	if filePaths, err = mgr.GetFiles(); err != nil {
		return
	}
	for _, filePath := range filePaths {
		var (
			pyp *PyParser
		)
		if pyp, err = NewPyParser(filePath, "s3.upload_file"); err != nil {
			return
		}
		if err = pyp.Process(); err != nil {
			return
		}
	}

	return
}
