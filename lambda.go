package aws_lambda_wrapper

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

func LambdaServe(router http.Handler) {
	adapter := httpadapter.New(router)
	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return adapter.Proxy(request)
	})
}
