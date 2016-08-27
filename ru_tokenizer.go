package shield

import (	
	"strconv"
    "github.com/legion-zver/shield/porterstemmers"
)


type ruTokenizer struct {
}

// NewRussianTokenizer - new Russian tokenizer
func NewRussianTokenizer() Tokenizer {
	return &ruTokenizer{}
}

func (t *ruTokenizer) Tokenize(text string) (words map[string]int64) {
    words = make(map[string]int64)
    pvStr := ""
    rusPS := porterstemmers.RussianPorterStemmer{}
    for _, w := range splitTokenRx.Split(text, -1) {
        if len(w) > 2 {
            if _, err := strconv.Atoi(w); err != nil {
                stem := rusPS.StemString(w)
                if len(stem) > 2 {
                    words[stem]++
                    if "" != pvStr {
                        bgWord := pvStr + "_" + stem							
                        words[bgWord]++
                    }
                    pvStr = stem
                }
            }
        }
    }
    return
}