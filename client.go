package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func getVersion(client *http.Client) {
	resp, err := client.Get("http://localhost:8080/version")
	if err != nil {
		log.Fatalf("Failed %v\n", err)
	}
	defer resp.Body.Close()
	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result["version"])
}

func decodeString(client *http.Client) {
	reqBody := `{"inputString": "U29tZSBCYXNlNjQgc3RyaW5n"}`
	resp, err := client.Post("http://localhost:8080/decode",
		"application/json", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		log.Fatalf("Failed %v\n", err)
	}
	defer resp.Body.Close()
	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result["outputString"])
}

func hardOp(client *http.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/hard-op", nil)
	if err != nil {
		log.Fatalln("EROERO")
	}

	resp, err := client.Do(req)
	if err != nil {
		if err == context.DeadlineExceeded {
			fmt.Println("false", err)
		}
	}
	defer resp.Body.Close()
	fmt.Println("true", resp.StatusCode)
}

func main() {
	client := &http.Client{}
	getVersion(client)
	decodeString(client)
	hardOp(client)
}
