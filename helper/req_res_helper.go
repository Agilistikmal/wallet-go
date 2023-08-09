package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequest(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&result)
	if err != nil {
		panic(err)
	}
}

func WriteToResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "Application/Json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	if err != nil {
		panic(err)
	}
}
