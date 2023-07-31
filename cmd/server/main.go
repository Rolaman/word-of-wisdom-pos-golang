package main

import (
	"log"
	"word-of-wisdom-pos/internal/common"
	"word-of-wisdom-pos/internal/server"
)

func main() {
	address := common.EnvVar("ADDRESS", "127.0.0.1:8001")
	store, err := server.NewStore(
		[]string{
			"The only true wisdom is in knowing you know nothing",
			"It does not matter how slowly you go as long as you do not stop",
			"In the middle of every difficulty lies opportunity",
		})
	if err != nil {
		log.Fatalf("can't init book store: %v", err)
	}
	provider, err := common.NewProvider(2)
	if err != nil {
		log.Fatalf("can't init pow provider: %v", err)
	}
	bookService := server.NewBook(store, provider)
	tcpServer := server.NewServer(address, bookService)
	if err := tcpServer.Start(); err != nil {
		log.Fatalf("can't start tcp tcpServer: %v", err)
	}
}
