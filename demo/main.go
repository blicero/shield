package main

import (     
    "github.com/legion-zver/shield"
)

func main() {    
     sh := shield.New(
        shield.NewRussianTokenizer(),
        shield.NewLevelDBStore("../db"),
    )
    sh.Reset()

    sh.Learn("good", "sunshine drugs love sex lobster sloth")
    sh.Learn("bad", "fear death horror government zombie god")

    c, _ := sh.Classify("sloths are so cute i love them")
    if c != "good" {
        panic(c)
    }

    c, _ = sh.Classify("i fear god and love the government")
    if c != "bad" {
        panic(c)
    }
}

