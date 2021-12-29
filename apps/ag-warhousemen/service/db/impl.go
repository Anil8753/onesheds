package db

import (
	"fmt"
)

type DB interface {
	// Put saves the data into the db
	Put(key string, value interface{}) error

	// Get retunrs the data for a key
	Get(key string) (interface{}, error)
}

type LevelDB struct {
	container map[string]interface{}
}

// Put saves the data into the db
func (s *LevelDB) Put(key string, value interface{}) error {
	s.container[key] = value
	return nil
}

// Get retunrs the data for a key
func (s *LevelDB) Get(key string) (interface{}, error) {

	val, ok := s.container[key]
	if !ok {
		return nil, fmt.Errorf("key '%s' not found", key)
	}

	return val, nil
}

func New() *LevelDB {
	return &LevelDB{container: make(map[string]interface{})}
}
