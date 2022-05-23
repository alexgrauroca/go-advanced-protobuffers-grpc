package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"platzi.com/go/grpc/testpb"
)

func main() {
	cc, err := grpc.Dial("localhost:5070", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	c := testpb.NewTestServiceClient(cc)
	DoUnary(c)
	DoClientStreaming(c)
}

func DoUnary(c testpb.TestServiceClient) {
	req := &testpb.GetTestRequest{
		Id: "t1",
	}

	res, err := c.GetTest(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GetTest: %v", err)
	}

	log.Printf("GetTest response: %v", res)
}

func DoClientStreaming(c testpb.TestServiceClient) {
	questions := []*testpb.Question{
		{
			Id:       "q8t1",
			Answer:   "lijasd",
			Question: "lijasdsad",
			TestId:   "t1",
		},
		{
			Id:       "q9t1",
			Answer:   "owiqeu",
			Question: "oiwque",
			TestId:   "t1",
		},
		{
			Id:       "q10t1",
			Answer:   "mnsd",
			Question: "okkwqwwwwww",
			TestId:   "t1",
		},
	}

	stream, err := c.SetQuestions(context.Background())

	if err != nil {
		log.Fatalf("Error while calling SetQuestions: %v", err)
	}

	for _, question := range questions {
		log.Println("Sending question: ", question.GetId())
		stream.Send(question)
	}

	msg, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while closing stream: %v", err)
	}

	log.Printf("response from server: %v", msg)
}
