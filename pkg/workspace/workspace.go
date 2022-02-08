package workspace

import (
	"fmt"

	"github.com/spf13/afero"
)

type Workspace struct {
	DiaryDir string
	IsLeap   bool
	Fs       *afero.Afero
}

// Create creates diary directory if it does not already exist.
func (ws *Workspace) Create() error {
	dirPath := fmt.Sprintf("diaries/%v", ws.DiaryDir)
	err := ws.Fs.MkdirAll(dirPath, 0755)
	if err != nil {
		return err
	} else {
		fmt.Printf("Create %v directory\n", ws.DiaryDir)
		return nil
	}
}
