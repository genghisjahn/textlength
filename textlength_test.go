package textlength

import "testing"

func TestOne(t *testing.T) {
	result, err := GetStringOfOnesInt(10)
	if err != nil {
		t.Error(err)
		return
	}
	correct := "one"
	if result != correct {
		t.Errorf("result is = %v, want %v.", result, correct)
	}
}
