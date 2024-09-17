package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ConfigResponse struct {
	Response string `json:"response"`
}

func main() {

	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://service-b:8081/internal/config")
		if err != nil {
			fmt.Println("Err: ", err)
			http.Error(w, "Failed to get config from service-b", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			http.Error(w, "Service-b returned an error", http.StatusInternalServerError)
			return
		}

		secretContent, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read config from service-b", http.StatusInternalServerError)
			return
		}

		// Map response to JSON
		config := ConfigResponse{
			Response: string(secretContent),
		}
		// Set Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")
		// Encode ConfigResponse to JSON and write to the response writer
		json.NewEncoder(w).Encode(config)
	})

	fmt.Println("Service A running on port 8080")
	http.ListenAndServe(":8080", nil)
}
