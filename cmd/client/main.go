package main

import (
	"log"
	"word-of-wisdom-pos/lib/client"
	"word-of-wisdom-pos/lib/common"
)

func main() {
	address := common.EnvVar("ADDRESS", "127.0.0.1:8001")
	tcpClient := client.NewClient(address)
	defer tcpClient.Close()
	err := tcpClient.Connect()
	if err != nil {
		log.Fatalf("can't connect to server: %v", err)
	}
	words, err := tcpClient.FetchWords()
	if err != nil {
		log.Fatalf("can't get words from server: %v", err)
	}
	log.Printf(words)
}
