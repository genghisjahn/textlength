package textlength

import "testing"

func TestOne(t *testing.T) {
	result, err := GetStringForOneDigit(1)
	if err != nil {
		t.Error(err)
		return
	}
	correct := "one"
	if result != correct {
		t.Errorf("result is = %v, want %v.", result, correct)
	}
}

func TestOneHundred(t *testing.T) {
	result, err := GetTextOfThreeDigitNumber(100, "")
	expected := "one hundred"
	if err != nil {
		t.Error(err)
		return
	}
	if result != expected {
		t.Errorf("Returned: %v\nExpected: %v", result, expected)
	}
}

func TestTextLengthItem(t *testing.T) {
	tli := TextLengthItem{"This text is forty characters long.", 20}
	if tli.Length() != 35 {
		t.Errorf("Length was %v, expected %v.", tli.Length(), 3)
	}
}
