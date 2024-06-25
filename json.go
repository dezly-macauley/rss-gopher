// NOTE: This is a helper function that will make it easier to send JSON
// reponses

package main

import (
	"encoding/json"
	"net/http"
    "log"
)

// code int : This is the status code that you respond with
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    dat, err := json.Marshal(payload)

    if err != nil {
        log.Printf("Failed to Marshal JSON response: %v", payload)
        w.WriteHeader(code)
        return
    }

    // NOTE: Adding a header to the response to show that you are reponding 
    //  with JSON data

    w.Header().Add("Content-Type", "application/json")

    // Write the status code
    w.WriteHeader(200)
    w.Write(dat)

}
