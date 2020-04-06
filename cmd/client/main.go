package main

import (
	"context"
	"log"
	"simple_string_api_grpc/proto"
	"time"

	"google.golang.org/grpc"
)

func main() {

	//Create a connection to the GRPC server, using the same port it is listening on.
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewManipulateServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.Reverse(ctx, &proto.Request{ToManipulate: "loafer"})

	log.Printf("%s\n", r.GetManipulated())

	r, err = client.NextInASCII(ctx, &proto.Request{ToManipulate: "loafer"})

	log.Printf("%s\n", r.GetManipulated())

	r, err = client.PreviousInASCII(ctx, &proto.Request{ToManipulate: "loafer"})

	log.Printf("%s\n", r.GetManipulated())

}
