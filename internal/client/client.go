package client

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"time"
	"word-of-wisdom-pos/internal/common"
)

type Client struct {
	address string
	con     net.Conn
}

func NewClient(address string) *Client {
	return &Client{
		address: address,
	}
}

func (c *Client) Connect() error {
	con, err := net.Dial("tcp", c.address)
	if err != nil {
		return err
	}
	if err := con.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
		log.Printf("can't set deadline: %v", err)
		return err
	}
	c.con = con
	challenge, diff, err := c.readChallenge()
	if err != nil {
		return fmt.Errorf("can't read challenge: %v", err)
	}
	solver, err := common.NewSolver(challenge, diff)
	if err != nil {
		return fmt.Errorf("can't init solver: %v", err)
	}
	nonce, err := solver.FindNonce()
	if err != nil {
		return err
	}
	err = c.sendSolution(challenge, nonce)
	if err != nil {
		return fmt.Errorf("can't send solution: %v", err)
	}
	return nil
}

func (c *Client) FetchWords() (string, error) {
	buf, err := ioutil.ReadAll(c.con)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func (c *Client) Close() {
	defer c.con.Close()
}

func (c *Client) readChallenge() ([]byte, byte, error) {
	buf := make([]byte, 9)
	if _, err := c.con.Read(buf); err != nil {
		return nil, 0, err
	}
	return buf[:8], buf[8], nil
}

func (c *Client) sendSolution(challenge []byte, nonce uint64) error {
	buffer := make([]byte, 8)
	binary.BigEndian.PutUint64(buffer, nonce)
	buffer = append(buffer, challenge...)
	if _, err := c.con.Write(buffer); err != nil {
		return err
	}
	return nil
}
