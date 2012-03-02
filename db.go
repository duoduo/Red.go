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
    return db.dict[db.hashKey(key)]
}

func (db *Db) hashKey(key []byte) string {
    return string(key)
}