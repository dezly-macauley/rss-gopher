// NOTE: This is a helper function that will make it easier to send JSON
// reponses

package main

import (
	"encoding/json"
	"net/http"
    "log"
)


func respondWithError(w http.ResponseWriter, code int, msg string) {

    if code > 499 {
        log.Println("Reponding with 5XX error:", msg)
    } 

    type errResponse struct {
        Error string `json:"error"`
    }
    respondWithJSON(w, code, errResponse {
        Error: msg,
    })

}

// code int : This is the status code that you respond with
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

    // Mashal the payload into a json object or a json string
    // and it will return it as bytes so that you can write it in a binary
    // format directly to the http response
    dat, err := json.Marshal(payload)

    if err != nil {
        log.Printf("Failed to Marshal JSON response: %v", payload)

        // If it fails, a head to the reponse will be written
        w.WriteHeader(500)
        return
    }

    // NOTE: Adding a header to the response to show that you are reponding 
    //  with JSON data

    w.Header().Add("Content-Type", "application/json")

    // Write the status code
    w.WriteHeader(code)
    w.Write(dat)

}
