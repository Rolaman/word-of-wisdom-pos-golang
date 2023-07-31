package server

import (
	"errors"
	"math/rand"
)

var (
	ErrEmptyQuotes = errors.New("store should have at least 1 quote")
)

type BookStore struct {
	quotes []string
}

func NewStore(quotes []string) (*BookStore, error) {
	if len(quotes) == 0 {
		return nil, ErrEmptyQuotes
	}
	return &BookStore{quotes: quotes}, nil
}

func (s *BookStore) GetRandomQuote() string {
	i := rand.Intn(len(s.quotes))
	return s.quotes[i]
}
