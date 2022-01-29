package workspace

import (
	"fmt"

	"github.com/spf13/afero"
)

type Workspace struct {
	diaryDir string
	isLeap   bool
	fs       *afero.Afero
}

// Create creates diary directory if it does not already exist.
func (ws *Workspace) Create() error {
	err := ws.fs.MkdirAll(ws.diaryDir, 0755)
	if err != nil {
		return err
	} else {
		fmt.Printf("Create %v directory", ws.diaryDir)
		return nil
	}
}
