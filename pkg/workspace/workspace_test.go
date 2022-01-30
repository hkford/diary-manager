package workspace

import (
	"testing"

	"github.com/spf13/afero"
)

func TestCreate(t *testing.T) {
	fs := afero.NewMemMapFs()
	ws := Workspace{
		DiaryDir: "2020",
		IsLeap:   true,
		Fs:       &afero.Afero{Fs: fs},
	}

	err := ws.Create()
	if err != nil {
		t.Fatal("failed Create workspace")
	}
}
