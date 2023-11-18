package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tahjib75/common"
)

func Handler(request events.APIGatewayProxyResponse) (events.APIGatewayProxyResponse, error) {

	db, err := common.ConnectToDB()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Database connection failed",
		}, nil
	}
	authors, err := db.GetAllAuthorWithRetry()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to fetch authors",
		}, nil
	}

	responseBody, _ := json.Marshal(authors)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(responseBody),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
