package main

import (
	"fmt"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)


func HandleRequest(_ context.Context) (string, error) {
	return fmt.Sprintf("Hello world from %s!", os.Getenv("MyEnvironment") ), nil
}

func main() {
	lambda.Start(HandleRequest)
}