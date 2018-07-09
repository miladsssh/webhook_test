package main

import (
	"fmt"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
	"log"
)


func HandleRequest(_ context.Context) (string, error) {
	return fmt.Sprintf("Hello world from %s!", os.Getenv("Myenvironment") ), nil
}

func main() {
	for _, e := range os.Environ() {
		log.Printf("Environmet variables: %s\n", e)
	}
	lambda.Start(HandleRequest)
}