package service

import (
	"go-challenge/config"
	"go-challenge/database"
	model "go-challenge/models"
	"log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

 
func Create(device model.Device) (model.Device, error) {

	dynamo := database.ConnectDynamo()

	_, err := dynamo.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(device.Id),
			},
			"deviceModel": {
				S: aws.String(device.DeviceModel),
			},
			"name": {
				S: aws.String(device.Name),
			},
			"note": {
				S: aws.String(device.Note),
			},
			"serial": {
				S: aws.String(device.Serial),
			},
		},
		TableName: aws.String(config.AppConfig.DbName),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			log.Println(aerr.Error())
		}
	} else {
		log.Println("Device is created successfully")
	}

	return model.Device{}, nil
}

 
 
func  GetById(id string) (device model.Device,e error) {

	dynamo := database.ConnectDynamo()
	result, err := dynamo.GetItem(&dynamodb.GetItemInput{
				Key: map[string]*dynamodb.AttributeValue{
					"id": {
						S: aws.String(id),
					},
				},
				TableName: aws.String(config.AppConfig.DbName),
			})

			if err != nil {
				if aerr, ok := err.(awserr.Error); ok {
					log.Println(aerr.Error())
				}
			}


		  	err = dynamodbattribute.UnmarshalMap(result.Item, &device)
 
			if err != nil {
				panic(err)
			}

	return device, nil
}
