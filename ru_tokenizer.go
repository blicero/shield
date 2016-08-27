package shield

import (	
	"strconv"
)


type ruTokenizer struct {
}

// NewRussianTokenizer - new Russian tokenizer
func NewRussianTokenizer() Tokenizer {
	return &ruTokenizer{}
}

func (t *ruTokenizer) Tokenize(text string) (words map[string]int64) {
    words = make(map[string]int64)
    for _, w := range splitTokenRx.Split(text, -1) {
        if len(w) > 2 {
            if _, err := strconv.Atoi(w); err != nil {

            }
        }
    }
    return
}