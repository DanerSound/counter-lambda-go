package main

import (
	//"log"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	//"github.com/aws/aws-lambda-go/lambda"
)

func lambda_handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	user := event.QueryStringParameters["user"]
	//log.Println(user)
	response := events.APIGatewayProxyResponse{
		Body: "hello " + user,
	}

	return response, nil
}

func main() {
	
		//Simulate an event
		event := events.APIGatewayProxyRequest{
			QueryStringParameters: map[string]string{"user": "DanerSound"},
		}

		// Invoke the Lambda function locally
		result, err := lambda_handler(event)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Response:", result.Body)
	
	//lambda.Start(lambda_handler)
}
