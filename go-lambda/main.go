package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response structure for the API response.
type Response struct {
	Message         string            `json:"message"`
	PathParameters  map[string]string `json:"path_parameters,omitempty"`
	QueryParameters map[string]string `json:"query_parameters,omitempty"`
}

// Handler handles API Gateway requests.
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Request received: %+v", request)

	// Extract query parameters
	queryParams := request.QueryStringParameters
	name := queryParams["name"]
	if name == "" {
		name = "World"
	}

	// Prepare the response
	response := Response{
		Message:         "Hello, " + name + "!",
		PathParameters:  request.PathParameters,
		QueryParameters: queryParams,
	}

	responseBody, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshaling response: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       `{"message": "Internal Server Error here"}`,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(responseBody),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
