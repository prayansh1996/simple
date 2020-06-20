package main

import (
	"context"
	"fmt"

	"github.com/prayansh1996/simple/proto/simple"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":11500", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Unable to reach server: " + err.Error())
		return
	}
	defer conn.Close()

	client := simple.NewSimpleServiceClient(conn)
	ctx := context.Background()

	res, err := client.SimpleRPC(ctx, &simple.SimpleRequest{SimpleString: "simple message"})
	if err != nil {
		fmt.Println("Unable to make request: " + err.Error())
		return
	}
	fmt.Println(res)
}
