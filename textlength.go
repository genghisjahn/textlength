package textlength

import "errors"
import "fmt"
import "strconv"

var singleValueErrorMessage = "Value supplied was %v.  Value must be %v %v."
var somethingWentWrongErrorMesasge = "Something went wrong in GetTextofThreeDigitNumber!"
var notImplementedYet = "Not implemented yet."

func GetTextOfThreeDigitNumber(i int, suffix string) (string, error) {
	result := ""
	maxLength := 1000
	stringInt := strconv.Itoa(i)
	//return stringInt, nil
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
			text, err := GetStringForOneDigit(num)
			result = fmt.Sprintf("%v hundred", text)
			return result, err
		}
	case 2:
		{
			return "", errors.New(notImplementedYet)
		}
	case 1:
		{
			num, _ := strconv.Atoi(string(stringInt[2]))
			text, err := GetStringForOneDigit(num)
			result = fmt.Sprintf("%v.", text)
			return result, err
		}
	}
	err_default := errors.New(somethingWentWrongErrorMesasge)
	return "", err_default
}

func GetStringForOneDigit(i int) (string, error) {
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
