package main

import (
	"context"
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
	//Not needed as it's already done above
	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("successfully connected and pinged the database")
	return nil
}

func main() {
	fmt.Println("Go REST API Course")

	err := Run()
	if err != nil {
		fmt.Println(err)
	}
}
