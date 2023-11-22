package main

import (
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tahjib75/common"
	"github.com/tahjib75/models"
	"github.com/tahjib75/utils"
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

	// err = common.ValidateAuthorSignup(signUpPayload)
	// if err != nil {
	// 	return events.APIGatewayProxyResponse{
	// 		StatusCode: 400,
	// 		Body:       "validation failed",
	// 	}, nil
	// }

	if signUpPayload.Password != signUpPayload.PasswordConfirm {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Password do not match",
		}, nil
	}

	hashpassword, err := utils.HashPassword(signUpPayload.Password)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to hash password",
		}, nil
	}
	// Create the author
	newAuthor := models.Author{
		UserName:  signUpPayload.UserName,
		Email:     signUpPayload.Email,
		Password:  hashpassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	token, refresh_token, err := utils.CreateToken(newAuthor.Email, newAuthor.Uid)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to create token",
		}, nil
	}

	newAuthor.Token = &token
	newAuthor.RefreshToken = &refresh_token

	db, err := common.ConnectToDB()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "database connection failed",
		}, nil
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
