package main

import (
	"log"
)

func main() {
	log.Println("starting server ...")

	data_source, err := init_db()

	if err != nil {
		log.Fatalf("Error when connecting with db, error:%v", err)
	}

	router, err := inject(data_source)

	if err != nil {
		log.Fatal(err)
	}

	router.Run(":8080")
}
