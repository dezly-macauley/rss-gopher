package main

import (
    "fmt"
    "os"
    "log"
    "net/http"

    // Load .env variables
    "github.com/joho/godotenv"

    "github.com/go-chi/chi"
    "github.com/go-chi/cors"
)


func main() {

    godotenv.Load(".env")
    
    
    // This will get the correct port number from the "PORT" key 
    // in the .env file
    portString := os.Getenv("PORT")

   if portString == "" {
        // This will exit the program immediately with an error code,
        // in case the "PORT" key is not in the .env file
        log.Fatal("No Port number was specified in environment configurations")
    }

    // This creates a new router object
    router := chi.NewRouter()

    // This adds a cor configuration:
    // This allows people to make HTTP requests from a browser

    // NOTE: This is fine for a local development but you definately want to 
    // tighten these up for a production environment

    router.Use(cors.Handler(cors.Options {
        AllowedOrigins: []string{"http://*", "http://*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"*"},
        ExposedHeaders: []string{"Link"},
        AllowCredentials: false,
        MaxAge: 300,
    }))


    // NOTE: Creating a new router

    v1Router := chi.NewRouter()

    // The handler readiness function is being connected to the 
    // "/healthz" path
    v1Router.HandleFunc("/healthz", handlerReadiness)

    router.Mount("/v1", v1Router)

    // NOTE: This server is going to be a JSON rest API
    // This means that all the requests and responses will have a JSON format

    // Next this router will be connected to a server
    srv := &http.Server{
        Handler: router,
        Addr: ":" + portString,
    }

    log.Printf("Sever starting on port %v", portString)
    err := srv.ListenAndServe()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Port:", portString)

}
