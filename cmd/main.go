package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"bdtbletask/internal/models"
)

type application struct {
	storage *models.StorageDB
}

const (
	storagePath = "./storage/storage.db"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <order_numbers>")
		return
	}

	orderNumbers := os.Args[1]
	orders := strings.Split(orderNumbers, ",")

	db, err := models.New(storagePath)
	if err != nil {
		log.Fatal(err)
		return
	}

	app := application{
		storage: &models.StorageDB{DB: db},
	}

	err = app.Manager(orders)
	if err != nil {
		log.Fatal(err)
	}
}
