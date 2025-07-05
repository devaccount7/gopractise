package main

import "parkinglot/router"

func main() {
	
	router := router.NewRouter()
	router.Run(":8080")
}