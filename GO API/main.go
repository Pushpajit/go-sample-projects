package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Pushpajit/go-api/router"
)

func main() {
	fmt.Println("Configuring the GO-API 🌎...")
	fmt.Println("Server is running at http://localhost:8000")

	r := router.Router()
	log.Fatal(http.ListenAndServe(":8000", r))

}
