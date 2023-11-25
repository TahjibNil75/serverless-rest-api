package common

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// NewInternalServerResponse returns a 500 Internal Server Error response
func NewInternalServerError(message string) events.APIGatewayProxyResponse {
	resp := map[string]interface{}{
		"detail": message,
	}
	val, _ := json.Marshal(resp)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(val),
	}
}

func NewUnauthorizedResponse(message string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 401,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `{"message": "` + message + `"}`,
	}
}

func NewSuccessResponse(message string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `{"message": "` + message + `"}`,
	}, nil
}

func NewBadRequestResponse(message string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 400,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `{"message": "` + message + `"}`,
	}, nil
}

func MethodNotAllowed(message string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 405,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `{"message": "` + message + `"}`,
	}, nil
}
