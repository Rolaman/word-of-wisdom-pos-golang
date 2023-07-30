package server

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"word-of-wisdom-pos/lib/common"
)

var (
	ErrInvalidSolution = errors.New("invalid solution")
	ErrDiffChallenge   = errors.New("got different challenge")
)

type BookService struct {
	store *BookStore
	pow   *common.PowProvider
}

func NewBook(store *BookStore, pow *common.PowProvider) *BookService {
	return &BookService{
		store: store,
		pow:   pow,
	}
}

func (b *BookService) HandleRequest(con net.Conn) error {
	if err := b.checkConnection(con); err != nil {
		return err
	}
	quote := b.store.GetRandomQuote()
	if _, err := con.Write([]byte(quote)); err != nil {
		return err
	}
	return nil
}

func (b *BookService) checkConnection(con net.Conn) error {
	originalChallenge := b.pow.GenerateChallenge()
	if _, err := con.Write(originalChallenge); err != nil {
		return err
	}
	buffer := make([]byte, 16)
	if _, err := io.ReadFull(con, buffer); err != nil {
		return err
	}
	nonce := binary.BigEndian.Uint64(buffer[:8])
	challenge := buffer[8:]
	if !bytes.Equal(originalChallenge[:8], challenge) {
		return ErrDiffChallenge
	}
	solution := common.NewSolution(challenge, nonce)
	result := b.pow.CheckSolution(solution)
	if !result {
		return ErrInvalidSolution
	}
	return nil
}
