package controllers

import (
	"encoding/json"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		handleGET(w, r)
	case http.MethodPost:
		handlePOST(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

// handleGET handles GET requests
func handleGET(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "This is a GET request")
	responseMessage := map[string]string{"message": "Get Api"}
	responseJSON, err := json.Marshal(responseMessage)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

// handlePOST handles POST requests
func handlePOST(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "This is a POST request")

	contentType := r.Header.Get("Content-Type")
	var responseMessage map[string]string
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if contentType == "application/json" {
		//Parsing Json Data
		var requestData map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}

		message, ok := requestData["message"].(string)
		if !ok {
			http.Error(w, "Invalid or missing 'message' field in JSON", http.StatusBadRequest)
			return
		}
		responseMessage = map[string]string{"message": "JSON data received successfully", "recivedMessage": message}
		responseJSON, err := json.Marshal(responseMessage)
		if err != nil {
			http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
			return
		}
		w.Write(responseJSON)
	} else if contentType == "application/x-www-form-urlencoded" {

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		// Access form values
		message := r.FormValue("message")

		responseMessage = map[string]string{"message": "data received successfully", "recivedMessage": message}
		responseJSON, err := json.Marshal(responseMessage)
		if err != nil {
			http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
			return
		}
		w.Write(responseJSON)

	} else {
		responseMessage = map[string]string{"message": "In valid content type"}
		responseJSON, err := json.Marshal(responseMessage)
		if err != nil {
			http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
			return
		}
		w.Write(responseJSON)
	}

}
