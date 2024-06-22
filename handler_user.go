package main

import (
    "encoding/json" // For JSON decoding
    "fmt"           // For formatting strings
    "net/http"      // For handling HTTP requests
    "time"          // For time.Now()
	
    

    // Replace "path/to/your/database/package" with the actual path to your database package
    
)
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter,r *http.Request){

	type parameters struct {
		Name string `json:"name"`
	}
	decoder:=json.NewDecoder(r.Body)
	params:=parameters{}
	err:=decoder.Decode(&params)

	if err!=nil{
		responseWithJSON(w,400,fmt.Sprintf("Error decoding request: %v",err))
		return 
	}

	user,err:=apiCfg.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID:uuid.New().String(),
		CreatedAt:time.Now().UTC(),
		UpdatedAt:time.Now().UTC(),
		Name:params.Name,
	})
	if err!=nil{
		respondWithJSON(w,400,fmt.Sprintf("Error creating user: %v",err))
		return 
	}
	responseWithJSON(w,200,struct{}{})
}