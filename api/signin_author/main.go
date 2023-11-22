package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tahjib75/common"
	"github.com/tahjib75/utils"
)

func Handler(request events.APIGatewayProxyResponse) (events.APIGatewayProxyResponse, error) {
	var signInPayLoad common.LoginValidator
	err := json.Unmarshal([]byte(request.Body), &signInPayLoad)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid Request body",
		}, nil
	}

	db, err := common.ConnectToDB()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error connecting to database",
		}, nil
	}

	authors, err := db.FindAuthor(signInPayLoad.Email)
	if err != nil {
		// Handle error
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error finding author",
		}, nil
	}

	// Assuming you only get one author, you can check the password
	author := authors[0]
	err = utils.VerifyPassword(author.Password, signInPayLoad.Password)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 401,
			Body:       "Invalid password",
		}, nil
	}

	// Password is valid
	token, err := utils.UpdateAllTokens(*author.Token, false)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error updating tokens",
		}, nil
	}

	responseBody := map[string]interface{}{
		"message": "Login successful",
		"token":   token,
	}

	responseJSON, err := json.Marshal(responseBody)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error marshalling response",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(responseJSON),
		Headers: map[string]string{
			"Set-Cookie": "Logged in=true",
		},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
