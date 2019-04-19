package fly

import (
	"encoding/json"
	"github.com/dgraph-io/badger"
	"log"
	"os"
	"os/user"
	"path"
)

const ServerPrefix = "SERVER_"
const DefaultValue = "DefaultValue"

type Server struct {
	Name string
	UserName string
	Host string
	Password string
	Port string
	Tag string
	Category string
}

func (server *Server) ToJson() []byte {
	if data, err:= json.Marshal(server); err == nil {
		return data
	}
	log.Fatal("Can't convert server to json.")
	return nil
}

func (server *Server) Merge(otherServer *Server) {
	if server.Port == "" {
		server.Port = otherServer.Port
	}
	if server.UserName == "" {
		server.UserName = otherServer.UserName
	}
	if server.Password == "" {
		server.Password = otherServer.Password
	}
}

func FromJson(data []byte) Server {
	var server Server
	if err := json.Unmarshal(data, &server); err == nil {
		return server
	}
	log.Fatal("Can't covert json to server.")
	return server
}


type DB struct {
	DB *badger.DB
}

func NewDB() *DB {
	return &DB{}
}

func (db *DB) Start() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	opts := badger.DefaultOptions
	opts.Dir = path.Join(usr.HomeDir, ".ssh", "fly")
	opts.ValueDir = opts.Dir
	db.DB, err = badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DB) close() {
	db.DB.Close()
}

func (db *DB) GetServer(serverName string) Server {
	var server Server
	if err := db.DB.View(func(txn *badger.Txn) error {
		var err error
		if item, err := txn.Get([]byte(ServerPrefix+serverName)); err == nil {
			if value, err := item.Value(); err == nil {
				server = FromJson(value)
			}
		}
		return err
	}); err != nil {
		log.Fatalf("Failed to view server %s", serverName)
	}
	return server
}

func (db *DB) GetServerList() []Server {
	var serverList []Server
	if err := db.DB.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte(ServerPrefix)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			key := item.Key()
			value, e := item.Value()
			if e == nil {
				serverList = append(serverList, FromJson(value))
			} else {
				log.Fatalf("Failed to get value for %s" , string(key))
			}
		}
		return nil
	}); err != nil {
		log.Fatal("Failed to view database.")
	}
	return serverList
}

func (db *DB) UpdateServer(server Server) {
	if err := db.DB.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(ServerPrefix+ server.Name), server.ToJson())
		return err
	}); err != nil {
		log.Fatal("Can not update server.")
	}
}

func (db *DB) UpdateDefault(server Server) {
	if err := db.DB.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(DefaultValue), server.ToJson())
		return err
	}); err != nil {
		log.Fatal("Can not update default.")
	}
}

func (db *DB) getDefault() Server {
	var server Server
	if err := db.DB.View(func(txn *badger.Txn) error {
		var err error
		if item, err := txn.Get([]byte(DefaultValue)); err == nil {
			if value, err := item.Value(); err == nil {
				server = FromJson(value)
			}
		}
		return err
	}); err != nil {
		log.Fatal("Failed to view default value")
	}
	return server
}
