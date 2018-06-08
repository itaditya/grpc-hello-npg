package main

import (
  "log"

  "golang.org/x/net/context"
  "google.golang.org/grpc"

  pb "helloworld/helloworldPb"
)

const (
  address = "localhost:50050"
)

func sayHello(client pb.GreeterClient, greeting *pb.HelloRequest) {
  resp, err := client.SayHello(context.Background(), greeting)
  if err != nil {
    log.Fatalf("Could not create Greeting: %v", err)
  }
  log.Printf("A new Greeting has been added with id: %s", resp.Message)
}

func main() {
  // Set up a connection to the gRPC server.
  conn, err := grpc.Dial(address, grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %v", err)
  }
  defer conn.Close()
  // Creates a new greeterClient
  client := pb.NewGreeterClient(conn)

  greeting := &pb.HelloRequest{
    Name: "Aditya",
  }

  // Create a new greeting
  sayHello(client, greeting)
}
