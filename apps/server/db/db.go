package db

type DB interface {
	// Put saves the data into the db
	Put(key string, value interface{}) error

	// Get retunrs the data for a key
	Get(key string) (interface{}, error)
}
