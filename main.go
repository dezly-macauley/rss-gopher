package main

import (
    "fmt"
    "os"
    "log"
    "net/http"

    // NOTE: This will automaticallly load variables from .env files into the
    // shell
    "github.com/joho/godotenv"

    // NOTE: This will be used to spin up the server
    "github.com/go-chi/chi"

    // NOTE: This will be used to handle the requests to the server
    "github.com/go-chi/cors"
)

func main() {

    // This is from the godotenv third party library. This will automaticallly
    // load the variable from the .env file into the shell, so that you don't
    // have to manually export each variable
    godotenv.Load(".env")
    
    // This will get the correct port number from the "PORT" key 
    // in the .env file
    portString := os.Getenv("PORT")

   if portString == "" {
        // This will exit the program immediately with an error code,
        // in case the "PORT" key is not in the .env file
        log.Fatal("No Port number was specified in environment configurations")
    }

    // NOTE: creates a new router object
    router := chi.NewRouter()

    // NOTE: This adds a cor configuration:
    // This allows people to make HTTP requests from a browser
    // This is fine for a local development but you definately want to 
    // tighten these up for a production environment

    router.Use(cors.Handler(cors.Options {
        AllowedOrigins: []string{"http://*", "http://*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"*"},
        ExposedHeaders: []string{"Link"},
        AllowCredentials: false,
        MaxAge: 300,
    }))

//=============================================================================

    v1Router := chi.NewRouter()

    // The handler readiness function is being connected to the 
    // "/healthz" path

    v1Router.Get("/healthz", handlerReadiness)

    v1Router.Get("/err", handlerErr)

    router.Mount("/v1", v1Router)

    // NOTE: So the full path for this request will be:
    // http://localhost:PORT/v1/healthz
    // Where PORT will be the one defined in the environment variables

    // For the error messages it will be 
    // http://localhost:8080/v1/err

//=============================================================================

    // NOTE: This server is going to be a JSON rest API
    // This means that all the requests and responses will have a JSON format

    // NOTE: This will connect the router to an http server
    // This is a server object

    srv := &http.Server{
        Handler: router,
        Addr: ":" + portString,
    }

    // Just a notification before the server starts listing for http requests
    log.Printf("Sever starting on port %v", portString)

    // NOTE: This line of code will block
    // The code will just stop here and start handling http requests

    err := srv.ListenAndServe()
    if err != nil {
        // If anything goes wrong in the process of handling those requests
        // an error will be logged and it will exit the program
        log.Fatal(err)
    }

    fmt.Println("Port:", portString)

}
