package main

import (   
    "fmt"  
    "github.com/legion-zver/shield"
)

func main() {
    store := shield.NewLevelDBStore("../db")    
    sh := shield.New(
        shield.NewRussianTokenizer(),
        store,
    )
    sh.Reset()    
	defer sh.Destroy()
	
	sh.Learn("a", "hello")
	sh.Learn("a", "sunshine")
	sh.Learn("a", "tree")
	sh.Learn("a", "water")
	sh.Learn("b", "iamb!")    
	sh.Forget("a", "hello tree")
	sh.Forget("a", "hello")
	
	m, err := store.ClassWordCounts("a", []string{
		"hello",
		"sunshine",
		"tree",
		"water",
	})
	if err != nil {
		fmt.Println(err)
        return
	}
	if r := fmt.Sprintf("%v", m); r != "map[hello:0 sunshine:1 tree:0 water:1]" {
		fmt.Println(r)
        return
	}

	m2, err := store.ClassWordCounts("b", []string{
		"hello",
		"iamb!",
	})
	if err != nil {
		fmt.Println(err)
        return
	}
	if r := fmt.Sprintf("%v", m2); r != "map[hello:0 iamb!:0]" {
		fmt.Println(r)        
        return
	}

	wc, err := store.TotalClassWordCounts()
	if err != nil {
		fmt.Println(err)
        return
	}
	if x := len(wc); x != 2 {
		fmt.Println(x)
        return
	}
	if x := wc["a"]; x != 2 {
		fmt.Println(x)
        return
	}
	if x := wc["b"]; x != 1 {
		fmt.Println(x)
        return
	}    
}

