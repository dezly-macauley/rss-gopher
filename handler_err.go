package main

import "net/http"

// NOTE: This is a very specific function signature
// This is how you define an HTTP handler in the way that the Go standard
// library expects

func handlerErr(w http.ResponseWriter, r *http.Request) {

    // struct{}{} means Marshal to an empty JSON object
    respondWithError(w, 400, "Something went wrong")
}


