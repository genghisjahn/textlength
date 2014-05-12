package textlength

import "errors"
import "fmt"

func GetStringOfOnesInt(i int) (string, error) {
	result := ""
	if i > 9 {
		err := errors.New(fmt.Sprintf("i Value supplied was %v.  i must be less than 10.", i))
		return result, err
	}
	if i == 1 {
		result = "one"
	}
	return result, nil
}
