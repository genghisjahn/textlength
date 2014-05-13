package textlength

import "fmt"

func BuildItems(numItems int) ([]TextLengthItem, error) {
	result := make([]TextLengthItem, numItems)
	for i := 1; i <= numItems; i++ {
		item := TextLengthItem{}
		item.Value = i
		item.Text, _ = GetTextForInt(i)
		result[i-1] = item
	}
	return result, nil
}

func ProcessText(text string) (string, error) {
	first_part := "This text is %v characters in length."
	result_template := "%v  %v"
	items, err := BuildItems(1000)
	textLength := len(text) + 2
	for _, item := range items {
		length_sentence := fmt.Sprintf(first_part, item.Text)
		if textLength+len(length_sentence) == item.Value {
			return fmt.Sprintf(result_template, text, length_sentence), nil
		}
	}
	return "", err
}
