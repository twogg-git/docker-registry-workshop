package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html><html><body><center>
		<img src="https://raw.githubusercontent.com/twogg-git/talks/master/resources/dockerhub-talk/hub-index-golang.png">
		<h1 style="color:blue">Baby steps with Docker Registry + Golang!</h1>
        </center></body></html>`
	fmt.Fprintf(w, html)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
