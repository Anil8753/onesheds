package db

import (
	"encoding/json"
	"os"
	"path"

	"github.com/syndtr/goleveldb/leveldb"
)

type LevelDB struct {
	db *leveldb.DB
}

// Put saves the data into the db
func (s *LevelDB) Put(key string, value interface{}) error {

	b, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return s.db.Put([]byte(key), b, nil)
}

// Get retunrs the data for a key
func (s *LevelDB) Get(key string) (interface{}, error) {

	b, err := s.db.Get([]byte(key), nil)
	if err != nil {
		return nil, err
	}

	var d interface{}
	err = json.Unmarshal(b, &d)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func NewLevelDB(collection string) *LevelDB {
	dataDir := os.Getenv("DATA_DIR")
	if dataDir == "" {
		panic("DATA_DIR is empty")
	}

	dbPath := path.Join(dataDir, "db", collection)

	// _, err := os.Stat(dbPath)
	// if os.IsNotExist(err) {
	// 	if err := os.Mkdir(dbPath, os.ModeDir); err != nil {
	// 		panic(err)
	// 	}
	// }

	dbInst, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		panic(err)
	}

	return &LevelDB{db: dbInst}
}
