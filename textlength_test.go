package textlength

import "testing"
import "errors"
import "fmt"

type TestItem struct {
	Num  int
	Text string
}

var ThreeOrLess = []TestItem{{1, "one"}, {10, "ten"}, {15, "fifteen"}, {16, "sixteen"},
	{42, "forty-two"}, {88, "eighty-eight"}, {100, "one hundred"}, {110, "one hundred ten"},
	{320, "three hundred twenty"}, {330, "three hundred thirty"}, {380, "three hundred eighty"},
	{562, "five hundred sixty-two"}, {814, "eight hundred fourteen"}, {900, "nine hundred"},
	{999, "nine hundred ninety-nine"}}

func TestThreeDigitOrLessNumbers(t *testing.T) {
	for _, ti := range ThreeOrLess {
		if err := runNumberTestAlt(ti.Num, ti.Text); err != nil {
			t.Error(err)
		}
	}
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

func TestOverOneBillion(t *testing.T) {
	runNumberTest(1000000000, "one billion", t)
	runNumberTest(999999999999, "nine hundred ninety-nine billion nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine", t)
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

func runNumberTestAlt(num int, text string) error {
	result, err := GetTextForInt(num)
	if err != nil {
		return err
	}
	if result != text {
		err = fmt.Errorf("\nGave: %v\nExpected: %v\nResult: %v", num, text, result)
		return err
	}
	return nil
}
