package singleton

import "fmt"

type Database struct {
}

var instance *Database

func (db *Database) GetInstance() *Database {

	if instance == nil {
		instance = &Database{}
		fmt.Println("Instance already created")
	} else {
		fmt.Println("Instance already exists")
	}
	return instance
}
