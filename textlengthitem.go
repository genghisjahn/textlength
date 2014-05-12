package textlength

type TextLengthItem struct {
	Text  string //length text add on.
	Value int    //Length of the supplied string to make the length text addon true.
}

func (t *TextLengthItem) Length() int {
	return len(t.Text)
}