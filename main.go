package main

import (
	"bdtbletask/internal/models"
	"log"
	"os"
)

type application struct {
	storage *models.StorageDB
}

const (
	storagePath = "/storage/storage.db"
)

func main() {
	pages := os.Args[1:]

	db, err := models.New(storagePath)
	if err != nil {
		log.Fatal(err)
		return
	}

	app := application{
		storage: &models.StorageDB{DB: db},
	}

	app.Manager(pages...)
}
