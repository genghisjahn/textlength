package textlength

import "errors"
import "fmt"
import "strconv"
import "math"
import "strings"

var singleValueErrorMessage = "\nValue supplied was %v.  Value must be %v %v."
var somethingWentWrongErrorMesasge = "\nSomething went wrong in GetTextofThreeDigitNumber!"
var notImplementedYet = "\nNot implemented yet."
var numberMustBeLessThanOneBillion = "\nValue supplied was %v.  Value must be less than one billion."

func GetTextForInt(i int) (string, error) {
	if i >= 1000000000 {
		return "", errors.New(fmt.Sprintf(numberMustBeLessThanOneBillion, i))
	}
	stringInt := strconv.Itoa(i)
	stringLen := len(stringInt)
	err := errors.New("")
	err = nil
	result := ""
	if stringLen > 0 && stringLen <= 3 {
		result, err = getTextOfThreeDigitNumber(i, "")
	}
	if stringLen > 3 && stringLen <= 6 {
		hundredsDigits := int(math.Mod(float64(i), float64(1000)))
		thousandsDigits := int(i / 1000)
		hundredsResult, _ := getTextOfThreeDigitNumber(hundredsDigits, "")
		thousandsResult, _ := getTextOfThreeDigitNumber(thousandsDigits, "thousand")
		result = fmt.Sprintf("%v %v", thousandsResult, hundredsResult)
	}
	if stringLen > 6 && stringLen <= 9 {
		thousandsDigits := int(math.Mod(float64(i), float64(1000000)))
		millionsDigits := int(i / 1000000)
		thousandsResult, _ := GetTextForInt(thousandsDigits)
		millionsResult, _ := getTextOfThreeDigitNumber(millionsDigits, "million")
		result = fmt.Sprintf("%v %v", millionsResult, thousandsResult)
	}
	return strings.TrimSpace(result), err
}

func getTextOfThreeDigitNumber(i int, suffix string) (string, error) {
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
			text, err := getTextForOneDigit(num)
			twoDigits := int(math.Mod(float64(i), float64(100)))
			text2, _ := getTextForTwoDigitNum(twoDigits)
			if twoDigits > 0 {
				result = fmt.Sprintf("%v hundred %v", text, text2)
			} else {
				result = fmt.Sprintf("%v hundred", text)
			}
			if suffix != "" {
				result = fmt.Sprintf("%v %v", result, suffix)
			}
			return result, err
		}
	case 2:
		{
			num, _ := strconv.Atoi(stringInt)
			text, err := getTextForTwoDigitNum(num)
			result = fmt.Sprintf("%v", text)
			if suffix != "" {
				result = fmt.Sprintf("%v %v", result, suffix)
			}
			return result, err
		}
	case 1:
		{
			num, _ := strconv.Atoi(stringInt)
			text, err := getTextForOneDigit(num)
			result = fmt.Sprintf("%v", text)
			if suffix != "" {
				result = fmt.Sprintf("%v %v", result, suffix)
			}
			return result, err
		}
	}
	err_default := errors.New(somethingWentWrongErrorMesasge)
	return "", err_default
}

func getTextForOneDigit(i int) (string, error) {
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
func getTextForTwoDigitNum(i int) (string, error) {
	result := ""
	tensPlace := ""
	if i < 16 {
		if i < 10 {
			result, _ = getTextForOneDigit(i)
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
	if i >= 20 {
		tensPlace = "twenty"
	}
	if i >= 30 {
		tensPlace = "thirty"
	}
	if i >= 40 {
		tensPlace = "forty"
	}
	if i >= 50 {
		tensPlace = "fifty"
	}
	if i >= 60 {
		tensPlace = "sixty"
	}
	if i >= 70 {
		tensPlace = "seventy"
	}
	if i >= 80 {
		tensPlace = "eighty"
	}
	if i >= 90 {
		tensPlace = "ninety"
	}
	onesDigit, _ := getTextForOneDigit(int(math.Mod(float64(i), float64(10))))
	if tensPlace != "teen" && onesDigit != "" {
		result = fmt.Sprintf("%v-%v", tensPlace, onesDigit)
	} else {
		result = fmt.Sprintf("%v%v", onesDigit, tensPlace)
	}

	return result, nil
}
