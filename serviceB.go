package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
Hauptfunktion
*/

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/v1").Subrouter()
	api.HandleFunc("/test/{userID}", homePage).Methods(http.MethodGet)

	r.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", r))
}
func homePage(w http.ResponseWriter, r *http.Request) {
	reqParameters := mux.Vars(r)
	var uri string
	uri = "http://localhost:8080/v1/user/" + reqParameters["userID"]
	client := &http.Client{}
	req, _ := http.NewRequest("GET", uri, nil)
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(body))
}
