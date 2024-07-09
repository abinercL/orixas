package routes

import (
	"api_test/controler"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", controler.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
