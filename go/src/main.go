package main

import (
	"fmt"
	"log"
	"net/http"
)

func greeting(message string) string {
    return fmt.Sprintf("<b>%s</b>", message)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, greeting("Code.education Rocks!") )
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
