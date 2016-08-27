package shield

import (
    "strconv"
    "github.com/syndtr/goleveldb/leveldb"
    "github.com/syndtr/goleveldb/leveldb/util"
)

// LevelDBStore store
type LevelDBStore struct {
    filePath string
    leveldb *leveldb.DB
}


// NewLevelDBStore - new redis store
func NewLevelDBStore(pathToDB string) Store {
	return &LevelDBStore{
		filePath: pathToDB,
	}    
}

func (ls *LevelDBStore) conn() (*leveldb.DB, error) {
    if ls.leveldb == nil {
        db, err := leveldb.OpenFile(ls.filePath, nil)
        if err != nil {
            return nil, err
        }
        ls.leveldb = db
    }    
    return ls.leveldb, nil
}

// Classes - list classes
func (ls *LevelDBStore) Classes() (a []string, err error) {
	db, err := ls.conn()
	if err != nil {
		return
	}
    iter := db.NewIterator(util.BytesPrefix([]byte("cs-")), nil)
	for iter.Next() {
        name := string(iter.Key())
        if len(name) > 3 {
            a = append(a, name[3:])
        }
    }
	return
}

// AddClass - add class
func (ls *LevelDBStore) AddClass(class string) (err error) {
	db, err := ls.conn()
	if err != nil {
		return
	}
	if class == "" {
		panic("invalid class: " + class)
	}
	err = db.Put([]byte("cs-"+class), []byte("1"), nil)
    return
}

// ClassWordCounts - count words by class
func (ls *LevelDBStore) ClassWordCounts(class string, words []string) (mc map[string]int64, err error) {
    db, err := ls.conn()
	if err != nil {
		return
	}
    mc = make(map[string]int64)
    for _, w := range words {
        key := "cl-"+class+"-"+w
        val, _ := db.Get([]byte(key), nil)
        if val != nil {
            count, e := strconv.ParseInt(string(val),10,64)
            if e == nil {
                mc[w] = count
            } else {
                mc[w] = 0
            }
        } else {
            mc[w] = 0
        }
    }
    return
}

// IncrementClassWordCounts - increment vals
func (ls *LevelDBStore) IncrementClassWordCounts(m map[string]map[string]int64) (err error) {
    db, err := ls.conn()
	if err != nil {
		return
	}
	type tuple struct {
		word string
		d    int64
	}    
    for class, words := range m {
		for word, d := range words {
            key := "cl-"+class+"-"+word           			
            appendIntValue(db, key, d)        
            appendIntValue(db, "sum-"+class, d)            
        }
    }    
    return
}

// TotalClassWordCounts - count words in classes
func (ls *LevelDBStore) TotalClassWordCounts() (m map[string]int64, err error) {
	db, err := ls.conn()
	if err != nil {
		return
	}
    m = make(map[string]int64)
    iter := db.NewIterator(util.BytesPrefix([]byte("sum-")), nil)
    for iter.Next() {
        name := string(iter.Key())
        if len(name) > 4 {
            count, e := strconv.ParseInt(string(iter.Value()),10,64)
            if e == nil {
                m[name[4:]] = count
            } else {
                m[name[4:]] = 0
            }
        }
    }
    return
}

// Reset - clean store
func (ls *LevelDBStore) Reset() (err error) {
	db, err := ls.conn()
	if err != nil {
		return
	}
    //batch := new(leveldb.Batch)            
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
        db.Delete(iter.Key(), nil)
    }
    //db.Write(batch, nil)
	return
}

// Close - close connection
func (ls *LevelDBStore) Close() error {
	if ls.leveldb != nil {
		err := ls.leveldb.Close()
		if err != nil {
			return err
		}
		ls.leveldb = nil
	}
	return nil
}

func appendIntValue(db *leveldb.DB, key string, d int64) {
    if db != nil {        
        val, e := db.Get([]byte(key), nil)
        if e != nil {     
            if d > 0 {           
                db.Put([]byte(key), []byte(strconv.FormatInt(d,10)), nil)
            }
        } else {
            v, e := strconv.ParseInt(string(val),10,64)
            if e != nil {
                if d > 0 {
                    db.Put([]byte(key), []byte(strconv.FormatInt(d,10)), nil)
                }
            } else {                
                if (v+d) < 0 {
                    d = v * -1
                }
                db.Put([]byte(key), []byte(strconv.FormatInt(d,10)), nil)
            }
        }
    }
}