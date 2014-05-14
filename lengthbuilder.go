package textlength

import "fmt"
import "errors"

func BuildItems(numItems int) ([]TextLengthItem, error) {
	result := make([]TextLengthItem, numItems)
	for i := 1; i <= numItems; i++ {
		item := TextLengthItem{}
		item.Value = i
		item.Text, _ = GetTextForInt(item.Value)
		result[i-1] = item
	}
	return result, nil
}

func ProcessText(text string, items []TextLengthItem) (string, error) {
	result:="Couldn't make this one work.  Sometimes the numbers just won't add up."
	first_part := "This text is %v characters in length."
	result_template := "%v  %v"
	textLength := len(text) + 2
	for _, item := range items {
		length_sentence := fmt.Sprintf(first_part, item.Text)
		if textLength+len(length_sentence) == item.Value {
			return fmt.Sprintf(result_template, text, length_sentence), nil
		}
	}
	err := errors.New(fmt.Sprintf("ProcessText for {%v} went wrong.", text))
	return result, err
}
