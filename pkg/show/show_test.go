package show

import (
	"mydiary/pkg/initialize"
	"mydiary/pkg/workspace"
	"testing"

	"github.com/spf13/afero"
)

func TestValidateInput(t *testing.T) {
	var result error
	_, result = ValidateInput(2000000)
	if result == nil {
		t.Errorf("ValidateInput failed at 2000000 %v", result)
	}
	_, result = ValidateInput(200000000)
	if result == nil {
		t.Errorf("ValidateInput failed at 200000000 %v", result)
	}
	_, result = ValidateInput(20160000)
	if result == nil {
		t.Errorf("ValidateInput failed at 20160000 %v", result)
	}
	_, result = ValidateInput(20170000)
	if result == nil {
		t.Errorf("ValidateInput failed at 20170000 %v", result)
	}
	_, result = ValidateInput(20171300)
	if result == nil {
		t.Errorf("ValidateInput failed at 20171300 %v", result)
	}
	_, result = ValidateInput(20170100)
	if result == nil {
		t.Errorf("ValidateInput failed at 20170100 %v", result)
	}
	_, result = ValidateInput(20170132)
	if result == nil {
		t.Errorf("ValidateInput failed at 20170131 %v", result)
	}
	_, result = ValidateInput(20200230)
	if result == nil {
		t.Errorf("ValidateInput failed at 20200230 %v", result)
	}
	date, result := ValidateInput(20200229)
	expected := Date{2020, 2, 29}
	if result != nil {
		t.Errorf("ValidateInput failed at 20200229 %v", result)
	}
	if date != expected {
		t.Errorf("ValidateInput expects %v but %v", expected, date)
	}
}

func TestOpenDiary(t *testing.T) {
	var result bool
	var date Date
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
	initialize.WriteMonthTemplate(ws, int64(1))
	date = Date{2020, 1, 29}
	result = IsDiaryFileExists(ws, date)
	if result != true {
		t.Errorf("diaries/2020/202001.txt should exist")
	}
	date = Date{2020, 3, 1}
	result = IsDiaryFileExists(ws, date)
	if result != false {
		t.Errorf("diaries/2020/202003.txt should not exist")
	}
	date = Date{2021, 2, 4}
	result = IsDiaryFileExists(ws, date)
	if result != false {
		t.Errorf("diaries/2021/202102.txt should not exist")
	}
}
