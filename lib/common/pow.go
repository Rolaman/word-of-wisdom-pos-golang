package common

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"time"
)

type PowProvider struct {
	difficulty uint8
}

type PowSolution struct {
	challenge []uint8
	nonce     uint64
}

type PowSolver struct {
	challenge  []uint8
	difficulty uint8
}

func NewProvider(difficulty uint8) (*PowProvider, error) {
	if difficulty < 0 || difficulty > 8 {
		return nil, fmt.Errorf("invalid difuculty: %v", difficulty)
	}
	return &PowProvider{difficulty: difficulty}, nil
}

func NewSolution(challenge []uint8, nonce uint64) PowSolution {
	return PowSolution{
		challenge: challenge,
		nonce:     nonce,
	}
}

func NewSolver(challenge []uint8, difficulty uint8) (*PowSolver, error) {
	if difficulty < 0 || difficulty > 8 {
		return nil, fmt.Errorf("invalid difuculty: %v", difficulty)
	}
	return &PowSolver{
		challenge:  challenge,
		difficulty: difficulty,
	}, nil
}

func (p *PowProvider) GenerateChallenge() []byte {
	now := time.Now()
	millis := now.UnixMilli()
	challengeBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(challengeBytes, uint64(millis))
	return challengeBytes
}

func (p *PowProvider) CheckSolution(solution PowSolution) bool {
	return checkSolution(solution.nonce, solution.challenge, p.difficulty)
}

func (s *PowSolver) FindNonce() (uint64, error) {
	for nonce := uint64(0); nonce <= ^uint64(0); nonce++ {
		if checkSolution(nonce, s.challenge, s.difficulty) {
			log.Printf("Found nonce %d", nonce)
			return nonce, nil
		}
	}
	return 0, fmt.Errorf("can't find solution for challenge: %v with difficulty %v", s.challenge, s.difficulty)
}

func checkSolution(nonce uint64, challenge []uint8, difficulty uint8) bool {
	nonceBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(nonceBytes, nonce)
	vector := append(nonceBytes, challenge...)
	hash := sha256.Sum256(vector)
	zeros := make([]byte, difficulty)
	if bytes.HasPrefix(hash[:], zeros) {
		return true
	}
	return false
}
