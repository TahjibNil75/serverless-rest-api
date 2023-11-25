package main

import (
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tahjib75/common"
	"github.com/tahjib75/middleware"
	"github.com/tahjib75/models"
)

func CreateArticle(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// check if user is logged in
	if request.RequestContext.Authorizer == nil {
		return common.NewUnauthorizedResponse("User is not logged in"), nil
	}

	// Parse the request body
	var articlePayload common.ArticleValidator
	err := json.Unmarshal([]byte(request.Body), &articlePayload)
	if err != nil {
		return common.NewBadRequestResponse("Invalid request body")
	}

	// Set the author id in the user logged in
	authorID, ok := request.RequestContext.Authorizer["Uid"].(uint)
	if !ok {
		return common.NewInternalServerError("User ID not found in RequestContext Authorizer"), nil
	}

	// Create a new article
	newArticle := models.Article{
		Name:      articlePayload.Name,
		Content:   articlePayload.Content,
		AuthorID:  int(authorID),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db, err := common.ConnectToDB()
	if err != nil {
		return common.NewInternalServerError("Could not connect to database"), nil
	}

	// save the article
	_, err = db.SaveArticle(&newArticle)
	if err != nil {
		return common.NewInternalServerError("Could not save article"), nil
	}

	// Return a success Response
	return common.NewSuccessResponse("article Posted Successfully")

}

func main() {
	createArticleHandler := middleware.AuthMiddleware(CreateArticle)

	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		if request.HTTPMethod == "POST" {
			return createArticleHandler(request)
		}
		return common.MethodNotAllowed("Method Not Allowed")
	})
}
