package main

import (
	"drumkit.com/interview/src/gateway"
	"drumkit.com/interview/src/handler"
	"drumkit.com/interview/src/service"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	gw := gateway.NewTurvoAPIGateway()
	svc := service.NewLoadService(gw)
	h := &handler.CreateLoadsHandler{Service: svc}
	lambda.Start(h.CreateLoadsHandlerLambda)
}
