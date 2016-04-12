package main

import (
	"log"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func madin() {
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8080", nil)
	// fmt.Println("Starting")
	log.Println("fw")
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.ListenAndServe(":3000", nil)
	// log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/Users/ramik/go-projects/FamilyScheduleServer/public"))))
}
