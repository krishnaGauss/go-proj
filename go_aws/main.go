package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
	age  int    `json:"age"`
}

type MyResponse struct {
	Message string `json:"answer:"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d years old", event.Name, event.age)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
