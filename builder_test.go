package textlength

import "testing"

func TestBuild1000(t *testing.T) {
	expected := "one thousand"
	items, err := BuildItems(1000)
	if err != nil {
		t.Error(err)
	}
	if items[len(items)-1].Text != expected {
		t.Errorf("Value: %v Expected: %v", items[len(items)-1].Text, expected)
	}
}

func TestBuild1000000(t *testing.T) {
	expected := "one million"
	items, err := BuildItems(1000000)
	if err != nil {
		t.Error(err)
	}
	if items[len(items)-1].Text != expected {
		t.Errorf("Value: %v Expected: %v", items[len(items)-1].Text, expected)
	}
}

func TestHello(t *testing.T) {
	expected := "Hello.  This text is fifty-two characters in length."
	text := "Hello."
	items, err := BuildItems(10000)
	if err != nil {
		text_with_length, _ := ProcessText(text, items)
		if expected != text_with_length {
			t.Errorf("Value: %v Expected: %v", text_with_length, expected)
		}
	}
	t.Error(err)
}

func TestHelloIsItMe(t *testing.T) {
	expected := "Hello.  Is it me you're looking for?  This text is eighty-four characters in length."
	text := "Hello.  Is it me you're looking for?"
	items, err := BuildItems(10000)
	if err != nil {
		text_with_length, _ := ProcessText(text, items)
		if expected != text_with_length {
			t.Errorf("Value: %v Expected: %v", text_with_length, expected)
		}
	}
	t.Error(err)
}
