package main

import (
	"context"
	"net"
	"simple_string_api_grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterManipulateServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}

func (s *server) Reverse(ctx context.Context, request *proto.Request) (*proto.Response, error) {

	var reversedString string

	//Load the variable with the word passed by the client.
	originalString := request.GetToManipulate()

	for i := range originalString {

		//The next line reverses the string by using slices building up reversedString.
		//For a 6 character string:
		//	Iteration 1 of the loop = originalString[5:6]
		//	Iteration 2 of the loop = originalString[4:5]
		//	Iteration 3 of the loop = originalString[3:4]
		//	Iteration 4 of the loop = originalString[2:3]
		//	Iteration 5 of the loop = originalString[1:2]
		//	Iteration 6 of the loop = originalString[0:1]
		reversedString = reversedString + originalString[len(originalString)-1-i:len(originalString)-i]
	}

	return &proto.Response{Manipulated: reversedString}, nil

}

func (s *server) NextInASCII(ctx context.Context, request *proto.Request) (*proto.Response, error) {

	var upscaledString string

	//Load the byte array with the word passed by the client.
	//Convert the word passed in the URL to an ASCII array
	byteArray := []byte(request.GetToManipulate())

	//Traverse through string, character by character.
	for i := range byteArray {

		upscaledString = upscaledString + string(byteArray[i]+1)

	}

	return &proto.Response{Manipulated: upscaledString}, nil

}

func (s *server) PreviousInASCII(ctx context.Context, request *proto.Request) (*proto.Response, error) {

	var downscaledString string

	//Load the byte array with the word passed by the client.
	//Convert the word passed in the URL to an ASCII array
	byteArray := []byte(request.GetToManipulate())

	//Traverse through string, character by character.
	for i := range byteArray {

		downscaledString = downscaledString + string(byteArray[i]-1)

	}

	return &proto.Response{Manipulated: downscaledString}, nil

}
