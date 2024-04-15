package main

import (
	pb "SPADE/spadeProto"
	"SPADE/utils"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/big"
	"time"
)

type User struct {
	id    int
	alpha *big.Int
}

func main() {
	addr := fmt.Sprintf("localhost:%d", utils.Port)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewCuratorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	q, g, _, err := getPublicParams(c.GetPublicParams(ctx, &pb.PublicParamsReq{}))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	utils.HexPrintBigInt("q", q)
	utils.HexPrintBigInt("g", g)
}

func getPublicParams(res *pb.PublicParamsRes, err error) (*big.Int, *big.Int, []*big.Int, error) {
	if err != nil {
		return nil, nil, nil, err
	}

	// Unmarshal q, g
	q := new(big.Int)
	q.SetBytes(res.Q)

	g := new(big.Int)
	g.SetBytes(res.G)

	// Unmarshal mpk (slice of big.Int)
	mpk := make([]*big.Int, 0, len(res.Mpk))
	for _, mpkBytes := range res.Mpk {
		temp := new(big.Int)
		temp.SetBytes(mpkBytes)
		mpk = append(mpk, temp)
	}

	return q, g, mpk, nil
}
