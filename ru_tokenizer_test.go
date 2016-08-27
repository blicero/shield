package shield

import (
	"fmt"
	"testing"
)

func TestRusTokenize(t *testing.T) {
	tokenizer := NewRussianTokenizer()
	text := "привет    ктото идет впереди кто    же это может быть  \t  жеесть! привет"
	m := tokenizer.Tokenize(text)
	x := fmt.Sprintf("%v", m)	
}