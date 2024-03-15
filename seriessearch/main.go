package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cw25/fuzzbun-api"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response := events.APIGatewayProxyResponse{
		Body:       "Hi there",
		StatusCode: 200,
	}

	fmt.Println("AKJSNKJSKJJKASKJN")
	fmt.Println(response)
	return "{ 'body': 'HI' }", nil
}

func main() {
	lambda.Start(handler)
}
