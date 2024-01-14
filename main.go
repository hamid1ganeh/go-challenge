package main

import (
	"context"
	"go-challenge/config"
	"go-challenge/database"
	route "go-challenge/router"
	Utility "go-challenge/utility"
	"log"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var echoLambda *echoadapter.EchoLambda

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	
	err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	database.CreateTable()

	server := echo.New()
	server.Validator = &Utility.CustomValidator{Validator: validator.New()}
	
	err = route.SetRoute(server)
	if err != nil {
		log.Fatal(err)
	}
	echoLambda = echoadapter.New(server)

	//in order run as locall without gateway
	// err = server.Start(":" + config.AppConfig.ServerPort)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	lambda.Start(handler)

}
