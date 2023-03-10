package main

import (
	"fmt"
	"net/http"

	"github.com/reotch/go-webapp/pkg/handlers"
)

const portNum = ":8080"

// main is the main function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNum))
	_ = http.ListenAndServe(portNum, nil)
}
