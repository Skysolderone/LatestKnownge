package server

import (
	"bufio"
	"net"
	"strings"

	"v1/pkg/pow"
	"v1/pkg/wisdom"
)

type Server struct {
	pow    pow.PoW
	wisdom wisdom.Wisdom
}

func NewServer(pow pow.PoWImpl, wisdom wisdom.WisdomImpl) *Server {
	return &Server{pow: pow, wisdom: wisdom}
}

func (s *Server) Handle(conn net.Conn) {
	defer conn.Close()
	seed, prefix := s.pow.GenerateChallege()
	_, _ = conn.Write([]byte("Solve PoW:" + seed + "with prefix" + prefix + "\n"))
	reader := bufio.NewReader(conn)
	proof, _ := reader.ReadString('\n')
	proof = strings.TrimSpace(proof)
	if !s.pow.VerifyPow(seed, proof) {
		quote := s.wisdom.GetWisdom()
		_, _ = conn.Write([]byte("Here is your wisdom:" + quote + "\n"))

	} else {
		_, _ = conn.Write([]byte("Invalid PoW!\n"))
	}
}
