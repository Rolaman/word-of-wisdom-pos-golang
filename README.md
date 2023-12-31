# Word of Wisdom Service with Proof Of Work DDOS Protection

## Task

**Design and implement a “Word of Wisdom” tcp server**

- TCP server should be protected from DDOS attacks with the Proof of Work, the challenge-response protocol should be used.
- The choice of the POW algorithm should be explained
- After Prof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes
- Docker file should be provided both for the server and for the client that solves the POW challenge

## Notes

- Libraries like zerolog or zap can be used for nicer logging. Skipped because of testing task goals

## How to run

Run server
```
docker-compose up -d server
// You can set env ADDRESS. Default - 0.0.0.0:8001
```

Run client
```
docker-compose up -d client
// You can set env ADDRESS. Default - server:8001
```

## Implementation

- The client establishes a new TCP connection with the server
- The server generates a new challenge(time-based) with a difficulty level of 2
- The server sends the message to the client
- The client parses the challenge and the difficulty
- The client calculates the right nonce to solve it
- The client sends the solution with the challenge
- If nonce is valid, the server sent a quote from a "word of wisdom" book
- If nonce is invalid, the server closes the stream

## PoW Algorithm

**I used the Hashcash with SHA256 hashing PoW algorithm to set DDoS protection**

- It is quick to implement, and many libraries provide SHA256 hashing
- It is easy to check the solution
- It doesn't consume much memory, so even clients with limited memory can use this service
- In this and many cases sha256 is sufficient to protect from DDoS attacks
- Using Hashcash is a common approach to limit spam and provide DDOS protection

**Tests and benchmarks can be found in the common/pow_test.rs file**

| Difficulty(leading zeros) | Average time to solve |
|---------------------------|-----------------------|
| 1                         | 21.71 µs              |
| 2                         | 4.27 ms               |
| 3                         | 2.00 s                |

## Problems and possible solutions

- TCP connection keep alive while waiting for solution from clients. It can be the reason of server stop.

Solution: Close current tcp connection to free server resources, GRPC protocol can be used

- Connection multi opening with one challenge in case of closing connection

Solution: Store session of each client, delete challenge after usage
