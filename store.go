package main

import (
	"errors"

	"github.com/peterbourgon/diskv"
)

var d *diskv.Diskv

// Initialize start store
func Initialize() {
	flatTransform := func(s string) []string { return []string{} }
	d = diskv.New(diskv.Options{
		BasePath:     "data",
		Transform:    flatTransform,
		CacheSizeMax: 1024 * 1024,
	})
}

// Set set datastore value
func Set(key string, value []byte) error {
	if d.Has(key) {
		return errors.New("value already exists")
	}
	errWrite := d.Write(key, value)
	if errWrite != nil {
		return errWrite
	}
	return nil
}

// Get datastore value
func Get(key string) ([]byte, error) {
	return d.Read(key)
}
