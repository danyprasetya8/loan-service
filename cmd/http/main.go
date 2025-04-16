package main

import "loan-service/app/http"

func main() {
	server := http.NewServer()
	server.Run()
}
