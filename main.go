package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("starting server ...")

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

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
