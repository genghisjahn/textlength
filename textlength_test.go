package textlength

import "testing"
import "errors"
import "fmt"

func TestThreeDigitOrLessNumbers(t *testing.T) {
	runNumberTest(1, "one", t)
	runNumberTest(10, "ten", t)
	runNumberTest(15, "fifteen", t)
	runNumberTest(16, "sixteen", t)
	runNumberTest(42, "forty-two", t)
	runNumberTest(88, "eighty-eight", t)
	runNumberTest(100, "one hundred", t)
	runNumberTest(110, "one hundred ten", t)
	runNumberTest(320, "three hundred twenty", t)
	runNumberTest(330, "three hundred thirty", t)
	runNumberTest(380, "three hundred eighty", t)
	runNumberTest(562, "five hundred sixty-two", t)
	runNumberTest(814, "eight hundred fourteen", t)
	runNumberTest(900, "nine hundred", t)
	runNumberTest(999, "nine hundred ninety-nine", t)
}

func TestFourToSixDigitNumbers(t *testing.T) {
	runNumberTest(2001, "two thousand one", t)
	runNumberTest(2010, "two thousand ten", t)
	runNumberTest(2012, "two thousand twelve", t)
	runNumberTest(2562, "two thousand five hundred sixty-two", t)
	runNumberTest(9000, "nine thousand", t)
	runNumberTest(9999, "nine thousand nine hundred ninety-nine", t)
	runNumberTest(12001, "twelve thousand one", t)
	runNumberTest(22010, "twenty-two thousand ten", t)
	runNumberTest(32012, "thirty-two thousand twelve", t)
	runNumberTest(42562, "forty-two thousand five hundred sixty-two", t)
	runNumberTest(59000, "fifty-nine thousand", t)
	runNumberTest(69999, "sixty-nine thousand nine hundred ninety-nine", t)
	runNumberTest(938469, "nine hundred thirty-eight thousand four hundred sixty-nine", t)
	runNumberTest(957395, "nine hundred fifty-seven thousand three hundred ninety-five", t)
	runNumberTest(791124, "seven hundred ninety-one thousand one hundred twenty-four", t)
	runNumberTest(642886, "six hundred forty-two thousand eight hundred eighty-six", t)
	runNumberTest(559621, "five hundred fifty-nine thousand six hundred twenty-one", t)
	runNumberTest(492319, "four hundred ninety-two thousand three hundred nineteen", t)
	runNumberTest(999999, "nine hundred ninety-nine thousand nine hundred ninety-nine", t)
}

func TestSevenToNineDigitNumbers(t *testing.T) {
	runNumberTest(2000000, "two million", t)
	runNumberTest(2032012, "two million thirty-two thousand twelve", t)
	runNumberTest(1938469, "one million nine hundred thirty-eight thousand four hundred sixty-nine", t)
	runNumberTest(62032776, "sixty-two million thirty-two thousand seven hundred seventy-six", t)
	runNumberTest(72132888, "seventy-two million one hundred thirty-two thousand eight hundred eighty-eight", t)
	runNumberTest(800000000, "eight hundred million", t)
	runNumberTest(813015010, "eight hundred thirteen million fifteen thousand ten", t)
	runNumberTest(999999999, "nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine", t)
}
func TestOneBillionOrMoreShouldFail(t *testing.T) {
	runNumberTest(1000000000, "one billion", t)
}

func runNumberTest(num int, text string, t *testing.T) (string, error) {
	result, err := GetTextForInt(num)
	if err != nil {
		t.Error(err)
	} else {
		if result != text {
			err = errors.New(fmt.Sprintf("\nGave: %v\nExpected: {%v}\nResult: {%v}", num, text, result))
			t.Error(err)
		}
	}
	return result, nil
}

func TestTextLengthItem(t *testing.T) {
	tli := TextLengthItem{"This text is forty characters long.", 20}
	if tli.Length() != 35 {
		t.Errorf("Length was %v, expected %v.", tli.Length(), 3)
	}
}
