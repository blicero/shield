package shield

import (
	"fmt"
	"testing"	
)

func TestRusTokenize(t *testing.T) {
	tokenizer := NewRussianTokenizer()
	text := "красивая красивее"
	m := tokenizer.Tokenize(text)
	fmt.Println(fmt.Sprintf("%v", m))
}