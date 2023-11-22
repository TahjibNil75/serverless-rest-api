package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tahjib75/common"
	"github.com/tahjib75/models"
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

	allAuthors := make([]models.GetAuthors, len(authors))
	for i, author := range authors {
		allAuthors[i] = models.GetAuthors{
			UserName: author.UserName,
			Email:    author.Email,
		}
	}

	responseBody, _ := json.Marshal(allAuthors)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(responseBody),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
