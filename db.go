package main

import(
    "fmt"
)

type Db struct {
    dict map[string] []byte
}

func NewDb() *Db {
    db := new(Db)
    db.dict = make(map[string] []byte, 100)
    return db
}

func (db *Db) Set(key []byte, value []byte) {
    db.dict[db.hashKey(key)] = value
}

func (db *Db) Get(key []byte) []byte {
    fmt.Printf("\n\n\ndb.dict: ", db.dict)
    data := db.dict[db.hashKey(key)]
    return data
}

func (db *Db) Delete(key []byte) {
    db.dict[db.hashKey(key)] = nil, false
}

func (db *Db) hashKey(key []byte) string {
    return string(key)
}