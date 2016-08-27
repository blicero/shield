package shield

import (
    "fmt"
    "testing"
)

func TestLevelDB(t *testing.T) {
    store := NewLevelDBStore("./db")
    defer store.Close()

    db, err := store.(*LevelDBStore).conn()
    if err != nil {
        t.Fatal(err)
    }    
    err = db.Put([]byte("test-value"), []byte("2024"), nil)
    if err != nil {
        t.Fatal(err)
    }
    val, err := db.Get([]byte("test-value"), nil)
    if err != nil {
        t.Fatal(err)
    }
    fmt.Println(string(val))      
}


func TestLevelDBLast(t *testing.T) {
    store := NewLevelDBStore("./db")    
    defer store.Close()

    db, err := store.(*LevelDBStore).conn()
    if err != nil {
        t.Fatal(err)
    }
    val, err := db.Get([]byte("test-value"), nil)
    if err != nil {
        t.Fatal(err)
    }
    fmt.Println(string(val))    
}
