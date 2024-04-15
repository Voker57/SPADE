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

const numUser = 10
const maxVecS = 1000

type Curator struct {
	q           *big.Int
	g           *big.Int
	sks         []*big.Int
	pks         []*big.Int
	regKeys     []*big.Int
	ciphertexts [][][]*big.Int
	spade       *SPADE.SPADE
}

type server struct {
	pb.UnimplementedCuratorServer
}

func NewCurator() *Curator {
	return &Curator{
		q:           nil,
		g:           nil,
		sks:         nil,
		pks:         nil,
		regKeys:     nil,
		ciphertexts: nil,
		spade:       nil,
	}
}

func (s *server) GetPublicParams(ctx context.Context, in *pb.PublicParamsReq) (*pb.PublicParamsRes, error) {
	log.Printf("CURATOR >>> Received GetPublicParams req!")

	cur := NewCurator()
	// generate q, q = (2 ^ 128) + 1
	q := new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil)
	q.Add(q, big.NewInt(1))
	// generate g
	g := SPADE.RandomElementInZMod(q)
	cur.q = q
	cur.g = g
	spd := SPADE.NewSpade(q, g)
	cur.sks, cur.pks = spd.Setup(numUser, maxVecS)
	cur.regKeys = make([]*big.Int, numUser)
	cur.ciphertexts = make([][][]*big.Int, numUser)
	cur.spade = spd

	// print q, g for debug
	utils.HexPrintBigInt("q", cur.q)
	utils.HexPrintBigInt("g", cur.g)

	qBytes := cur.q.Bytes()
	gBytes := cur.g.Bytes()
	mpkBytes := make([][]byte, 0, len(cur.pks)) // Pre-allocate for efficiency
	for _, pk := range cur.pks {
		mpkBytes = append(mpkBytes, pk.Bytes())
	}

	return &pb.PublicParamsRes{
		Q:   qBytes,
		G:   gBytes,
		Mpk: mpkBytes,
	}, nil
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

}
