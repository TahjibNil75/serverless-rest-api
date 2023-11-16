package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tahjib75/common"
	"github.com/tahjib75/models"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Parse the request body
	var signUpPayload common.AuthorSignupValidator
	err := json.Unmarshal([]byte(request.Body), &signUpPayload)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid request body",
		}, nil
	}

	err = common.ValidateAuthorSignup(signUpPayload)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "validation failed",
		}, nil
	}

	db, err := common.ConnectToDB()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "database connection failed",
		}, nil
	}

	// Create the author
	newAuthor := models.Author{
		UserName: signUpPayload.UserName,
		Email:    signUpPayload.Email,
	}
	createdAuthor, err := db.CreateAuthor(&newAuthor)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to create author",
		}, nil
	}

	reponseBody, _ := json.Marshal(createdAuthor)
	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Body:       string(reponseBody),
	}, nil

}

func main() {
	lambda.Start(Handler)
}
