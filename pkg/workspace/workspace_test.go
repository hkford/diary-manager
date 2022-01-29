package workspace

import (
	"testing"

	"github.com/spf13/afero"
)

func TestCreate(t *testing.T) {
	fs := afero.NewMemMapFs()
	ws := Workspace{
		diaryDir: "2020",
		isLeap:   true,
		fs:       &afero.Afero{Fs: fs},
	}

	err := ws.Create()
	if err != nil {
		t.Fatal("failed Create workspace")
	}
}
