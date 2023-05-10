package main

import (
	"go-patterns/di"
	"go-patterns/singleton"
	"log"
	"sync"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error occurred while initializing env variables: %s", err)
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		di.Init()
		wg.Done()
	}()

	go func() {
		singleton.Init()
		wg.Done()
	}()

	wg.Wait()
}
