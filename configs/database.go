package configs

import (
	// "fmt"
	"log"
	// "os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// func ConnectToDB() {
// 	// Lee las variables de entorno
// 	dbHost := os.Getenv("DB_HOST")
// 	dbPort := os.Getenv("DB_PORT")
// 	dbUser := os.Getenv("DB_USER")
// 	dbPassword := os.Getenv("DB_PASSWORD")
// 	dbName := os.Getenv("DB_NAME")
// 	dbCharset := os.Getenv("DB_CHARSET")
// 	dbLoc := os.Getenv("DB_LOC")

// 	// Construye la cadena de conexión DSN
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
//     dbUser, dbPassword, dbHost, dbPort, dbName, dbCharset, dbLoc)

// 	var err error

// 	// Conéctate a la base de datos utilizando la cadena de conexión DSN
// 	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatalf("Failed to connect DB: %v", err)
// 	}
// }

func ConnectToDB() {
	var err error

	dsn := "angelina:2162012@tcp(db:3306)/go_crud_docker?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect DB")
	}
}
