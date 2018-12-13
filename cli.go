package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/micahlagrange/grpc-demo/protodemo"
)

func main() {
	// Create instance of Person from the generated go
	// This requires that you first run
	//     mkdir protos
	//     protoc -I=. --go_out=protos person.proto
	// See the imports above
	p := protodemo.Person{
		Name:     "Bob",
		Age:      12,
		Birthday: 2135246.0,
		Book:     "This is a book about something cool!",
	}

	// Marshall the person to bytes
	x, err := proto.Marshal(&p)
	if err != nil {
		log.Fatalf("Could not marshal: %v", err)
	}

	// Prove it was encoded to bytes
	fmt.Println(x)

	// unmarshal the byte slice into y
	var y protodemo.Person
	err = proto.Unmarshal(x, &y)
	if err != nil {
		log.Fatalf("Could not unmarshal: %v", err)
	}

	// Prove that the bytes were decoded as an instance of protodemo.Person by calling Book from it:
	fmt.Println(y.Book)
}
