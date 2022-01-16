package db

import "fmt"

type Mock struct {
	container map[string]interface{}
}

// Put saves the data into the db
func (s *Mock) Put(key string, value interface{}) error {
	s.container[key] = value
	return nil
}

// Get retunrs the data for a key
func (s *Mock) Get(key string) (interface{}, error) {

	val, ok := s.container[key]
	if !ok {
		return nil, fmt.Errorf("key '%s' not found", key)
	}

	return val, nil
}

func NewMock() *Mock {
	return &Mock{container: make(map[string]interface{})}
}
