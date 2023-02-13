package main

import (
	"log"
)

func main() {
	svc := NewCatFactService("https://catfact.ninja/fact")
	svc = NewLoggingService(svc)

	apiService := NewApiServer(svc)
	log.Fatal(apiService.start(":3000"))
}
