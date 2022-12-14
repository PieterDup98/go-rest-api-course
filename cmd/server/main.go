package main

import (
	"context"
	"fmt"

	"github.com/PieterDup98/go-rest-api-course/db"
	"github.com/PieterDup98/go-rest-api-course/internal/comment"
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

	//Feature logic
	cmtService := comment.NewService(db)

	cmtCreate, errCreate := cmtService.CreateComment(context.Background(), comment.Comment{
		Slug:   "manual-test",
		Body:   "Hello World",
		Author: "Pieter",
	})
	if errCreate == nil {
		fmt.Println("inserted:", cmtCreate)
	}

	cmtCreate.Author = "John Wayne"
	cmtUpdate, errUpdate := cmtService.UpdateComment(context.Background(), cmtCreate)
	if errUpdate == nil {
		fmt.Println("updated:", cmtUpdate)
	}

	var staticUUId string = "28e7a8d7-3aee-4451-afba-4d04176b6f9b"
	cmtGet, errGet := cmtService.GetComment(context.Background(), staticUUId)
	if errGet == nil {
		fmt.Println("fetched:", cmtGet)
		cmtService.DeleteComment(context.Background(), staticUUId)
	} else {
		fmt.Println("couldn't fetch " + staticUUId + " skipping delete")
	}

	return nil
}

func main() {
	fmt.Println("Go REST API Course")

	err := Run()
	if err != nil {
		fmt.Println(err)
	}
}
