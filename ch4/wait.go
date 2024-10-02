package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	println("Hello World")
	err := WaitForServer("http://localhost:8080")
	fmt.Println(err)
}

func WaitForServer(url string) error {
	const timeout = 60 * time.Second
	start := time.Now()
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		fmt.Printf("time gap %v\n", time.Now().Sub(start))
		_, err := http.Get(url)
		if err == nil {
			return nil
		}
		log.Printf("Waiting for server to start... %d", tries)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
