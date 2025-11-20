package creational

import "sync"

// Singleton - Creational Pattern
// Ensures a class has only one instance and provides global access to it

type Database struct {
	connection string
}

var instance *Database
var once sync.Once

func GetDatabaseInstance() *Database {
	once.Do(func() {
		instance = &Database{connection: "Connected"}
	})
	return instance
}
