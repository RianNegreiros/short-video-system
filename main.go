package main

import (
	"github.com/RianNegreiros/short-video-system/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := models.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
