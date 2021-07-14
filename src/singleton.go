package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct{}

func (Database) createConnection() {
	fmt.Println("creating db connection")
	time.Sleep(3 * time.Second)
	fmt.Println("db connection created successfully\n")
}

var db *Database
var lock sync.Mutex

func GetDatabaseInstance() *Database {
	lock.Lock()
	defer lock.Unlock()

	if db == nil {
		fmt.Println("database does not exist, instanciating")
		db = &Database{}
		db.createConnection()
	} else {
		fmt.Println("database already exists, getting instance")
	}
	return db
}

func main() {
	nGoRoutines := 5
	var wg sync.WaitGroup
	wg.Add(nGoRoutines)

	for i := 0; i < nGoRoutines; i++ {
		go func() {
			defer wg.Done()
			GetDatabaseInstance()
		}()
	}

	wg.Wait()
}
