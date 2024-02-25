package main

import (
	"github.com/AngelinaSantana/Docker-/configs"
	"github.com/AngelinaSantana/Docker-/models"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.Person{})
}
