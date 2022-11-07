package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	request1()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Welcome to go web api")
	fmt.Println("Endpoint hit: homepage")
}

func aboutMe(w http.ResponseWriter, r *http.Request) {
	who := "CeliaSaumell"

	fmt.Fprintf(w, "A little about meeeeeeeeeee..")
	fmt.Println("endpoint hit ", who)

}

func request1() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutme", aboutMe)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
