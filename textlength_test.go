package textlength

import "testing"

func Test1(t *testing.T) {
	result, err := GetTextForOneDigit(1)
	if err != nil {
		t.Error(err)
		return
	}
	correct := "one"
	if result != correct {
		t.Errorf("result is = %v, want %v.", result, correct)
	}
}

func Test10(t *testing.T) {
	expected := "ten"
	result, err := GetTextForTwoDigitNum(10)
	if err != nil {
		t.Error(err)
		return
	}
	if result != expected {
		t.Errorf("Returned: %v\nExpected: %v", result, expected)
	}
}

func Test16(t *testing.T) {
	expected := "sixteen"
	result, err := GetTextForTwoDigitNum(16)
	if err != nil {
		t.Error(err)
		return
	}
	if result != expected {
		t.Errorf("Returned: %v\nExpected: %v", result, expected)
	}
}

func Test15(t *testing.T) {
	expected := "fifteen"
	result, err := GetTextForTwoDigitNum(15)
	if err != nil {
		t.Error(err)
		return
	}
	if result != expected {
		t.Errorf("Returned: %v\nExpected: %v", result, expected)
	}
}

func Test42(t *testing.T) {
	expected := "forty-two"
	result, err := GetTextForTwoDigitNum(42)
	if err != nil {
		t.Error(err)
		return
	}
	if result != expected {
		t.Errorf("Returned: %v\nExpected: %v", result, expected)
	}
}

func Test88(t *testing.T) {
	expected := "eighty-eight"
	result, err := GetTextForTwoDigitNum(88)
	if err != nil {
		t.Error(err)
		return
	}
	if result != expected {
		t.Errorf("Returned: %v\nExpected: %v", result, expected)
	}
}

func Test110(t *testing.T) {
	result, err := GetTextOfThreeDigitNumber(110, "")
	expected := "one hundred ten"
	if err != nil {
		t.Error(err)
		return
	}
	if result != expected {
		t.Errorf("Returned: %v\nExpected: %v", result, expected)
	}
}

func Test562(t *testing.T) {
	result, err := GetTextOfThreeDigitNumber(562, "")
	expected := "five hundred sixty-two"
	if err != nil {
		t.Error(err)
		return
	}
	if result != expected {
		t.Errorf("Returned: %v\nExpected: %v", result, expected)
	}
}

func Test814(t *testing.T) {
	result, err := GetTextOfThreeDigitNumber(814, "")
	expected := "eight hundred fourteen"
	if err != nil {
		t.Error(err)
		return
	}
	if result != expected {
		t.Errorf("Returned: %v\nExpected: %v", result, expected)
	}
}

func Test100(t *testing.T) {
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

func Test900(t *testing.T) {
	result, err := GetTextOfThreeDigitNumber(900, "")
	expected := "nine hundred"
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
