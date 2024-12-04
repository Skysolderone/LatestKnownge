package client

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"

	"v1/pkg/pow"
)

func Start(conn net.Conn) {
	reader := bufio.NewReader(conn)
	message, _ := reader.ReadString('\n')
	parts := strings.Fields(message)
	if len(parts) < 3 {
		fmt.Println("Invalid challenge format received")
		return
	}
	seed := parts[2]
	proof := solvePoW(seed)
	_, _ = conn.Write([]byte(proof + "\n"))
	response, _ := reader.ReadString('\n')
	fmt.Println("Server response:", response)
}

func solvePoW(seed string) string {
	var proof int
	powImpl := pow.PoWImpl{}

	for {
		proofstr := strconv.Itoa(proof)
		if powImpl.VerifyPow(seed, proofstr) {
			return proofstr
		}
		proof++
	}
}
