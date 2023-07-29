package server

import (
	"fmt"
	"math/rand"
)

type BookStore struct {
	quotes []string
}

func NewStore(quotes []string) (*BookStore, error) {
	if len(quotes) == 0 {
		return nil, fmt.Errorf("store should have at least 1 quote")
	}
	return &BookStore{quotes: quotes}, nil
}

func (s *BookStore) GetRandomQuote() string {
	i := rand.Intn(len(s.quotes))
	return s.quotes[i]
}
