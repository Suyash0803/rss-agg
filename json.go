package main

import(
	"encoding/json"
	"log"
	"net/http"
)

func respondwithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	responseWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func responseWithJSON(w http.ResponseWriter,code int, payload interface{}){
	response, err:=json.Marshal(payload)
	if err!=nil{
		log.Fatal(err)
		w.WriteHeader(500)
		return 
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(response)
}