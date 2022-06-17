package main

import (
	"UserServicePractice/DataUser"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var PORT = ":8080"

var  users = map[int]DataUser.User{
	1: {
		Id: 1,
		Username: "hehehe",
		Email: "Michael.edrick@gmail.com",
		Password: "hehehe123",
		Age: 17,
	},
	2: {
		Id: 2,
		Username: "hehehe",
		Email: "Michael.edrick@gmail.com",
		Password: "hehehe123",
		Age: 17,
	},
}

func main(){

	r := mux.NewRouter()
	r.HandleFunc("/employees", userHandler)
	r.HandleFunc("/employees/{Id}", userHandler)
	
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

	// http.HandleFunc("/employees/", userHandler)
	// fmt.Println("Application is listening on PORT", PORT)
	// http.ListenAndServe(PORT, nil)
}

func userHandler(w http.ResponseWriter, r *http.Request){
	param := mux.Vars(r)
	id := param["Id"]

	if r.Method == "GET"{
		if id != "" {
			fmt.Println(id)
			if idInt, err := strconv.Atoi(id); err != nil{
				return
			} else{
				jsonData, _ :=json.Marshal(users[idInt])
				w.Header().Add("Content-Type", "application/json")
				w.Write(jsonData)
			}
		} else {
			fmt.Println(id)
			fmt.Println("no param")
			//kudu convert ke slice lagi sebelum di json.marshal
			sliceUsers := []DataUser.User{
			}
			for _, value := range users {
				sliceUsers = append(sliceUsers, value)
			}
			jsonData, _ := json.Marshal(&sliceUsers)
			w.Header().Add("Content-Type", "application/json")
			w.Write(jsonData)
		}
	}
	
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
    	var newUsers DataUser.User
    	err := decoder.Decode(&newUsers)
		if err != nil {
			panic(err)
		}
		users[int(newUsers.Id)] = newUsers
		//users = append(users, newUsers)
		fmt.Println(users)
		sliceUsers := []DataUser.User{
		}
		for _, value := range users {
			sliceUsers = append(sliceUsers, value)
		}
		json.NewEncoder(w).Encode(sliceUsers)
		return
	} 
	if r.Method == "PUT" {
		decoder := json.NewDecoder(r.Body)
    	var newUsers DataUser.User
    	err := decoder.Decode(&newUsers)
		if err != nil {
			panic(err)
		}
		users[int(newUsers.Id)] = newUsers
		//users = append(users, newUsers)
		fmt.Println(users)
		sliceUsers := []DataUser.User{
		}
		for _, value := range users {
			sliceUsers = append(sliceUsers, value)
		}
		json.NewEncoder(w).Encode(sliceUsers)
		return
	}
	if r.Method == "DELETE" {
		if id != "" {
			if	index, err := strconv.Atoi(id); err != nil{
				return
			}else{
				delete(users, index)
			}
		}
	}
}