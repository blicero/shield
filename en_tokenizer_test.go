package shield

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	tokenizer := NewEnglishTokenizer()
	text := "lorem    ipsum able hello erik    can do hi there  \t  spaaace! lorem"
	m := tokenizer.Tokenize(text)	
	wait := map[string]int64{"lorem":2, "ipsum":1, "hello":1, "erik":1, "spaaace":1}
	if !compareMapSI(m, wait) {
		t.Fatal(m)
	}
}
