package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kiloMIA/Final_SRE/internal"
)

func main() {
    logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal("Could not create log file: ", err)
    }
    defer logFile.Close()

    logger := log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

    pool, err := internal.ConnectDB()
    if err != nil {
        logger.Fatalf("Unable to connect to database: %v\n", err)
    }
    defer pool.Close()

    router := internal.NewRouter(pool)

    logger.Println("Starting server on :8081")
    if err := http.ListenAndServe(":8081", router); err != nil {
        logger.Fatalf("Error starting server: %v\n", err)
    }
}
