package textlength

import "errors"
import "fmt"
import "strconv"
import "math"

var singleValueErrorMessage = "Value supplied was %v.  Value must be %v %v."
var somethingWentWrongErrorMesasge = "Something went wrong in GetTextofThreeDigitNumber!"
var notImplementedYet = "Not implemented yet."

func GetTextOfThreeDigitNumber(i int, suffix string) (string, error) {
	result := ""
	maxLength := 1000
	stringInt := strconv.Itoa(i)
	if i > maxLength {
		err := errors.New(fmt.Sprintf(singleValueErrorMessage, i, "less than", maxLength))
		return "", err
	}
	if i <= 0 {
		err := errors.New(fmt.Sprintf(singleValueErrorMessage, i, "greater than", 0))
		return "", err
	}
	switch len(stringInt) {
	case 3:
		{
			num, _ := strconv.Atoi(string(stringInt[0]))
			text, err := GetTextForOneDigit(num)
			twoDigits := int(math.Mod(float64(i), float64(100)))
			text2, _ := GetTextForTwoDigitNum(twoDigits)
			if twoDigits > 0 {
				result = fmt.Sprintf("%v hundred %v", text, text2)
			} else {
				result = fmt.Sprintf("%v hundred", text)
			}
			return result, err
		}
	case 2:
		{
			num, _ := strconv.Atoi(string(stringInt[1:]))
			text, err := GetTextForTwoDigitNum(num)
			result = fmt.Sprintf("%v", text)
			return result, err
		}
	case 1:
		{
			num, _ := strconv.Atoi(string(stringInt[2]))
			text, err := GetTextForOneDigit(num)
			result = fmt.Sprintf("%v", text)
			return result, err
		}
	}
	err_default := errors.New(somethingWentWrongErrorMesasge)
	return "", err_default
}

func GetTextForOneDigit(i int) (string, error) {
	result := ""
	if i > 9 {
		err := errors.New(fmt.Sprintf(singleValueErrorMessage, i, "less than", 10))
		return result, err
	}

	switch i {
	case 9:
		result = "nine"
	case 8:
		result = "eight"
	case 7:
		result = "seven"
	case 6:
		result = "six"
	case 5:
		result = "five"
	case 4:
		result = "four"
	case 3:
		result = "three"
	case 2:
		result = "two"
	case 1:
		result = "one"
	case 0:
		result = ""
	}

	return result, nil
}
func GetTextForTwoDigitNum(i int) (string, error) {
	result := ""
	tensPlace := ""
	if i < 16 {
		if i < 10 {
			result, _ = GetTextForOneDigit(i)
		}
		switch i {
		case 10:
			{
				result = "ten"
			}
		case 11:
			{
				result = "eleven"
			}
		case 12:
			{
				result = "twelve"
			}
		case 13:
			{
				result = "thirteen"
			}
		case 14:
			{
				result = "fourteen"
			}
		case 15:
			{
				result = "fifteen"
			}
		}
		return result, nil
	}
	if i > 15 {
		tensPlace = "teen"
	}
	if i > 20 {
		tensPlace = "twenty"
	}
	if i > 30 {
		tensPlace = "thirty"
	}
	if i > 40 {
		tensPlace = "forty"
	}
	if i > 50 {
		tensPlace = "fifty"
	}
	if i > 60 {
		tensPlace = "sixty"
	}
	if i > 70 {
		tensPlace = "seventy"
	}
	if i > 80 {
		tensPlace = "eighty"
	}
	if i > 90 {
		tensPlace = "niney"
	}
	onesDigit, _ := GetTextForOneDigit(int(math.Mod(float64(i), float64(10))))
	if tensPlace != "teen" {
		result = fmt.Sprintf("%v-%v", tensPlace, onesDigit)
	} else {
		result = fmt.Sprintf("%v%v", onesDigit, tensPlace)
	}

	return result, nil
}
