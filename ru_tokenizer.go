package shield

import (
    "fmt"
    "regexp"
    "strings"
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
    rusPS := porterstemmers.RussianPorterStemmer{}
    for _, w := range strings.Split(replaceWSpacesRx.ReplaceAllString(text, " ")," ") {
        if len(w) > 2 {
            if _, err := strconv.Atoi(w); err != nil {
                stem := rusPS.StemString(w)
                fmt.Println(w, " - ", stem)
                if len(stem) > 2 {
                    words[stem]++                    
                }
            }
        }
    }
    return
}

var replaceWSpacesRx = regexp.MustCompile(`[\s|\n]+`)