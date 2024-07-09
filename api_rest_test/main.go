package main

import (
	"api_test/routes"
	"fmt"
)

func main() {
	fmt.Println("Init server rest")
	routes.HandleRequest()
}
