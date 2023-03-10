package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(OnEC2StateChange)
}

func OnEC2StateChange(ctx context.Context) (err error) {
	fmt.Println("Hello, World!")
	return nil
}
