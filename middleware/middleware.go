package middleware

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/tahjib75/utils"
)

// func AuthMiddleware(next func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)) func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 		// Check if the user is logged in based on the request headers and cookies
// 		// For simplicity, I'm checking for a "Logged in" cookie here.
// 		if _, ok := request.Headers["Cookie"]; !ok || request.Headers["Cookie"] != "Logged in=true" {
// 			return events.APIGatewayProxyResponse{
// 				StatusCode: http.StatusUnauthorized,
// 				Body:       "Unauthorized: User not logged in",
// 			}, nil
// 		}

// 		// User is logged in, call the next handler in the chain.
// 		return next(ctx, request)
// 	}
// }

func AuthMiddleware(next func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)) func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		// Get the Authorization header from the request
		token := request.Headers["Authorization"]
		if token == "" {
			return events.APIGatewayProxyResponse{
				StatusCode: 400,
				Body:       "Authorization token missing",
			}, nil
		}

		// Validate the token
		email, uid, err := utils.ValidateToken(token)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 400,
				Body:       "Invalid Token",
			}, nil
		}

		// Convert the models, Author to map[string]interface{}
		authorMap := map[string]interface{}{
			"Email": email,
			"Uid":   uid,
		}

		// Add user information to the request context for later use
		request.RequestContext.Authorizer = authorMap

		// Call the next handler in the chain
		return next(request)
	}
}
