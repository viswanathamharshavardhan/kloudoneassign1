package main

import (
	"fmt"
	"log"
	"net/http"
)

func docker(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Kloudone, Harsha")

}

func routers() {
	http.HandleFunc("/", docker)
}

func main() {
	routers()
	log.Fatal(http.ListenAndServe(":8081", nil))

}
