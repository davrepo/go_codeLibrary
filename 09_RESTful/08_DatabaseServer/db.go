// db.go
package main

import "sync"

type DB struct {
	m sync.Map
}

func (db *DB) Get(key string) []byte {
	val, ok := db.m.Load(key)
	if !ok {
		return nil
	}

	// val.(type) is a type assertion expression
	// that asserts that val is of type []byte
	return val.([]byte)
}

func (db *DB) Set(key string, value []byte) {
	db.m.Store(key, value)
}
