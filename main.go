package main

import (
    "github.com/joho/godotenv"
    "shortner/routes"
    "shortner/db"
    "log"

)

func init() {
    // Log error if .env file does not exist
    if err := godotenv.Load(); err != nil {
        log.Printf("No .env file found")
    }
}

func main() { 
    db.GetConnection()
    routes.Initialize()
    
}


