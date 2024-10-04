// /home/krylon/go/src/github.com/blicero/shield/de_tokenizer_test.go
// -*- mode: go; coding: utf-8; -*-
// Created on 11. 10. 2022 by Benjamin Walkenhorst
// (c) 2022 Benjamin Walkenhorst
// Time-stamp: <2022-10-11 21:42:44 krylon>

package shield

import "testing"

func TestDeTokenize(t *testing.T) {
	var (
		tokenizer = NewGermanTokenizer()
		text      = `
Himmel und Erde waren geschaffen: das Meer wogte
in seinen Ufern, und die Fische spielten darin; in den
Lüften sangen beflügelt die Vögel; der Erdboden wimmelte
von Thieren.
`
		expect = map[string]int64{
			"himmel":     1,
			"erde":       1,
			"geschaffen": 1,
			"meer":       1,
			"wogte":      1,
			"ufern":      1,
			"fische":     1,
			"spielten":   1,
			"lüften":     1,
			"sangen":     1,
			"beflügelt":  1,
			"vögel":      1,
			"erdboden":   1,
			"wimmelte":   1,
			"thieren":    1,
		}
		result = tokenizer.Tokenize(text)
	)

	if !compareMapSI(result, expect) {
		t.Fatal(result)
	}
}
