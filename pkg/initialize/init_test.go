package initialize

import (
	"mydiary/pkg/workspace"
	"testing"

	"github.com/spf13/afero"
)

func TestGetDate(t *testing.T) {
	var result string
	result = GetDate(2022, 1, 28)
	if result != "Fri" {
		t.Fatal("failed GetDate 2022/01/28")
	}
	result = GetDate(2020, 3, 9)
	if result != "Mon" {
		t.Fatal("failed GetDate 2020/03/09")
	}
}

func TestGenerateDayFormat(t *testing.T) {
	var result string
	result = GenerateDayFormat(2022, 1, 1)
	if result != "2022,January,01,Sat\n\n" {
		t.Fatal("failed GenerateDayformat(2022,1,1)")
	}
}

func TestWriteMonthTemplate(t *testing.T) {
	fs := afero.NewMemMapFs()
	ws := workspace.Workspace{
		DiaryDir: "2020",
		IsLeap:   true,
		Fs:       &afero.Afero{Fs: fs},
	}

	err := ws.Create()
	if err != nil {
		t.Fatal("failed Create workspace")
	}
	WriteMonthTemplate(ws, int64(1))
	_, err = fs.Open("2020/202001.txt")
	if err != nil {
		t.Fatal("Failed to write monthly template")
	}
}
