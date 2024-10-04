package shield

import (
	"fmt"
	"testing"	
)

func TestRusTokenize(t *testing.T) {
	tokenizer := NewRussianTokenizer()
	text := "красивая красивее"
	m := tokenizer.Tokenize(text)
	x := fmt.Sprintf("%v", m)
	if x != `map[красив:2]` {
		t.Fatal(x)
	}
}