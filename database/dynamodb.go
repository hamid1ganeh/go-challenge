package database

import (
	"log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"go-challenge/config"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

var dynamo *dynamodb.DynamoDB


func ConnectDynamo() (db *dynamodb.DynamoDB) {
 
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.AppConfig.AWSREGION),
		Credentials: credentials.NewStaticCredentials(config.AppConfig.AwsAccesskey, config.AppConfig.AwsSecretKey, ""),
	})
	if err != nil {
		log.Fatalf("Failed to create AWS session: %v", err)
	}
	return dynamodb.New(sess)
}


func CreateTable() {

	dynamo = ConnectDynamo()

	_, err := dynamo.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
		TableName: aws.String(config.AppConfig.DbName),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {

			log.Println(aerr.Error())
		}
	}
}