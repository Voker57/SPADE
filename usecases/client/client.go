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
	"os"
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
	name := "greeting"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Setup(ctx, &pb.PublicParams{M: name})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.Msg)
}
