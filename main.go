package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Handler to get the last active timestamp for a given transaction ID
func getLastActiveTimestamp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	txid := vars["txid"]

	url := fmt.Sprintf("https://vayu.hornet.technology/api/tx/%s", txid)

	var responseData map[string]interface{}

	resp, err := http.Get(url)
	if err != nil {
		responseData = map[string]interface{}{
			"message": "failure",
			"error":   err.Error(),
		}
	} else {
		defer resp.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			responseData = map[string]interface{}{
				"message": "failure",
				"error":   err.Error(),
			}
		} else {

			var data map[string]interface{}
			err = json.Unmarshal(body, &data)
			if err != nil {
				responseData = map[string]interface{}{
					"message": "failure",
					"error":   err.Error(),
				}
			} else {

				var timestamp float64
				if timeValue, ok := data["time"]; ok {
					timestamp = timeValue.(float64)
				} else {
					responseData = map[string]interface{}{
						"message": "failure",
						"error":   "time field not found",
					}
				}

				if responseData == nil {

					humanReadableTime := time.Unix(int64(timestamp), 0).Format("2006-01-02 15:04:05")
					// timestamp = timestamp

					responseData = map[string]interface{}{
						"message":               "success",
						"last_active_timestamp": humanReadableTime,
					}
				}
			}
		}
	}

	responseJSON, err := json.MarshalIndent(responseData, "", "  ")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello! This is the home page.")
	w.Write([]byte("\n"))
	w.Write([]byte("To get the last active timestamp for a transaction ID, make a GET request to /api/tx/{txid}"))
	w.Write([]byte("\n"))
	w.Write([]byte("Example: /api/tx/1a2b3c4d5e6f7g8h9i0j"))

}
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/api/tx/{txid}", getLastActiveTimestamp).Methods("GET")

	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
