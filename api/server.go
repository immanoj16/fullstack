package api

import (
	"fmt"
	"log"
	"os"

	"github.com/immanoj16/fullstack/api/controllers"
	"github.com/immanoj16/fullstack/api/seed"
	"github.com/joho/godotenv"
)

// Run is the server root
func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
		return
	}

	fmt.Println("We are getting the env variables")

	server := controllers.Server{}

	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	server.Initialize(dbDriver, dbUser, dbPass, dbPort, dbHost, dbName)

	seed.Load(server.DB)
	server.Run(":8080")
}
