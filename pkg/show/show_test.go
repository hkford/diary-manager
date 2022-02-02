package show

import "testing"

func TestValidateInput(t *testing.T) {
	var result error
	result = ValidateInput(2000000)
	if result == nil {
		t.Errorf("ValidateInput failed at 2000000 %v", result)
	}
	result = ValidateInput(200000000)
	if result == nil {
		t.Errorf("ValidateInput failed at 200000000 %v", result)
	}
	result = ValidateInput(20160000)
	if result == nil {
		t.Errorf("ValidateInput failed at 20160000 %v", result)
	}
	result = ValidateInput(20170000)
	if result == nil {
		t.Errorf("ValidateInput failed at 20170000 %v", result)
	}
	result = ValidateInput(20171300)
	if result == nil {
		t.Errorf("ValidateInput failed at 20171300 %v", result)
	}
	result = ValidateInput(20170100)
	if result == nil {
		t.Errorf("ValidateInput failed at 20170100 %v", result)
	}
	result = ValidateInput(20170132)
	if result == nil {
		t.Errorf("ValidateInput failed at 20170131 %v", result)
	}
	result = ValidateInput(20200230)
	if result == nil {
		t.Errorf("ValidateInput failed at 20200230 %v", result)
	}
	result = ValidateInput(20200229)
	if result != nil {
		t.Errorf("ValidateInput failed at 20200229 %v", result)
	}
}
