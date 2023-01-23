package main

import (
	"fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
	"encoding/json"
)

type Response struct {

	Id   int `json:"id"`
	Uid  string `json:"uid"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Username string `json:"username"`

}




func main() {

	response, err := http.Get("https://random-data-api.coml/api/v2/users")
	if err!= nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    //fmt.Println(string(responseData))

	

    var responseObject Response
 	json.Unmarshal(responseData, &responseObject)

	fmt.Println("Id", responseObject.Id)
	fmt.Println("Uid" ,responseObject.Uid)
	fmt.Println("First Name",responseObject.FirstName)
	fmt.Println("Last Name",responseObject.LastName)
	fmt.Println("Email", responseObject.Email)
	fmt.Println("Username",responseObject.Username)
	
	
} 

