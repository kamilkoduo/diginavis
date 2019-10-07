package main

import (
	"github.com/kamilkoduo/diginavis/src/adder"
	"github.com/kamilkoduo/diginavis/src/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	api.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ": 8080")
	if err != nil {
		log.Fatal(err)
	}
	err = s.Serve(l);
	if err != nil {
		log.Fatal(err)
	}
}
