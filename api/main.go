package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	fmt.Printf("Body size = %d.\n", len(request.Body))

	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %s: %s\n", key, value)
	}

	// process the request against the routes


	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}

func main() {
	exitIfError(loggerCreate())
	exitIfError(versionPrint())
	exitIfError(configParse())
	exitIfError(dbConnect(5))
	exitIfError(migrate())
	exitIfError(smtpConfigure())
	exitIfError(smtpTemplatesLoad())
	exitIfError(oauthConfigure())
	exitIfError(markdownRendererCreate())
	exitIfError(emailNotificationPendingResetAll())
	exitIfError(emailNotificationBegin())
	exitIfError(sigintCleanupSetup())
	exitIfError(versionCheckStart())
	exitIfError(domainExportCleanupBegin())
	exitIfError(viewsCleanupBegin())
	exitIfError(ssoTokenCleanupBegin())

	//exitIfError(routesServe()) //this is replaced by lambda request handler
	// instead we need to just setup the routes
	exitIfError(routesSetup())

	// handle the request
	lambda.Start(handleRequest)
}
