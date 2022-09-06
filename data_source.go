package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dataSource struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}

func init_db() (*dataSource, error) {
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", db_host, db_user, db_pass, db_name, db_port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err == nil {
		seed(db)
	} else {
		log.Fatalf("Failed connection to postgres database, error: %v", err)
	}

	redis_host := os.Getenv("REDIS_HOST")
	redis_port := os.Getenv("REDIS_PORT")

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redis_host, redis_port),
		Password: "",
		DB:       0,
	})

	_, err = rdb.Ping(context.Background()).Result()

	if err != nil {
		log.Fatalf("Failed connection to redis database, error: %v", err)
	}

	return &dataSource{
		DB:          db,
		RedisClient: rdb,
	}, err
}
