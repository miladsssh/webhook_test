package main

import (
	"fmt"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
	"strings"
	"log"
)


func HandleRequest(_ context.Context) (string, error) {
	return fmt.Sprintf("Hello world from %s!", os.Getenv("MyEnvironment") ), nil
}

func main() {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		log.Printf("Environmet variables: %s\n", pair[0])
	}
	lambda.Start(HandleRequest)
}