package main

import (
	"SPADE"
	pb "SPADE/spadeProto"
	"SPADE/utils"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math/big"
	"net"
)

type curator struct {
	sks   []*big.Int
	pks   []*big.Int
	spade *SPADE.SPADE
}

type server struct {
	pb.UnimplementedCuratorServer
}

func NewCurator() *curator {
	return &curator{
		sks:   nil,
		pks:   nil,
		spade: nil,
	}
}

// Setup generates the secret and public keys when you want to initialize the system
func (cur curator) Setup(numUser, maxVecSize int) {
	spd := SPADE.NewSpade()
	cur.sks, cur.pks = spd.Setup(numUser, maxVecSize)
	cur.spade = spd
}

func (s *server) Setup(ctx context.Context, in *pb.PublicParams) (*pb.Empty, error) {
	log.Printf("Received: %v", in.GetM())
	//return &pb.HelloRep{Msg: "Hello, " + in.GetMsg()}, nil
	return &pb.Empty{}, nil
}

func main() {
	addr := fmt.Sprintf(":%d", utils.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	log.Printf("Server listening on port %d", lis.Addr())

	pb.RegisterCuratorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// setup spade
	cur := NewCurator()
	cur.Setup(1, 1000)

}
