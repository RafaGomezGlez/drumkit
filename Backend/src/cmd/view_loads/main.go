package viewloads

import (
	"drumkit.com/interview/src/gateway"
	"drumkit.com/interview/src/handler"
	"drumkit.com/interview/src/service"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	gw := gateway.NewTurvoAPIGateway()
	svc := service.NewLoadService(gw)
	h := &handler.ViewLoadsHandler{Service: svc}
	lambda.Start(h.ViewLoadsHandlerLambda) // Corrected function name to match the handler's method
}
