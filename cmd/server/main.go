package main

import (
	"fmt"

	"github.com/PieterDup98/go-rest-api-course/db"
)

// Run - is going to be responsible for
// the instantiation and startup of our
// go application
func Run() error {
	fmt.Println("starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to the database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	fmt.Println("successfully connected pinged and migrated the database")
	return nil
}

func main() {
	fmt.Println("Go REST API Course")

	err := Run()
	if err != nil {
		fmt.Println(err)
	}
}
