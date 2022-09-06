package main

import (
	"fmt"
	"log"
	"os"
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

	port := os.Getenv("PORT")

	router.Run(fmt.Sprintf(":%v", port))
}
