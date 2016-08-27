package main

import (
    "fmt"
    "github.com/legion-zver/shield/porterstemmers"
)

func main() {
    rusPS := porterstemmers.RussianPorterStemmer{}
	fmt.Println(rusPS.StemString("красивее"))
}