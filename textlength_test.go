package textlength

import "testing"
import "errors"
import "fmt"

func TestNumbers(t *testing.T) {
	runNumberTest(1, "one", t)
	runNumberTest(10, "ten", t)
	runNumberTest(15, "fifteen", t)
	runNumberTest(16, "sixteen", t)
	runNumberTest(42, "forty-two", t)
	runNumberTest(88, "eighty-eight", t)
	runNumberTest(100, "one hundred", t)
	runNumberTest(110, "one hundred ten", t)
	runNumberTest(320, "three hundred twenty", t)
	runNumberTest(562, "five hundred sixty-two", t)
	runNumberTest(814, "eight hundred fourteen", t)
	runNumberTest(900, "nine hundred", t)
}

func runNumberTest(num int, text string, t *testing.T) (string, error) {
	result, err := GetTextOfThreeDigitNumber(num, "")
	if err != nil {
		t.Error(err)
	}
	if result != text {
		err = errors.New(fmt.Sprintf("\nGave: %v\nExpected: %v\nResult: %v", num, text, result))
		t.Error(err)
	}
	return result, nil
}

func TestTextLengthItem(t *testing.T) {
	tli := TextLengthItem{"This text is forty characters long.", 20}
	if tli.Length() != 35 {
		t.Errorf("Length was %v, expected %v.", tli.Length(), 3)
	}
}
