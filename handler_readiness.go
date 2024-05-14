package main

import (
	"net/http"
)

func headerReadiness(w http.ResponseWriter,r *http.Request){
	responseWithJSON(w,200,struct{}{})
}