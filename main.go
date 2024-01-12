package main

import (
	"go-challenge/config"
	"go-challenge/database"
	route "go-challenge/router"
	Utility "go-challenge/utility"
	"log"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

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

	err = server.Start(":" + config.AppConfig.ServerPort)
	if err != nil {
		log.Fatal(err)
	}

}

//var dynamo *dynamodb.DynamoDB
//
//type Person struct {
//	Id      int
//	Name    string
//	Website string
//}
//
//const TABLE_NAME = "peoples"
//
//func init() {
//	dynamo = connectDynamo()
//}
//
//func connectDynamo() (db *dynamodb.DynamoDB) {
//	return dynamodb.New(session.Must(session.NewSession(&aws.Config{
//		Region: aws.String("us-east-1"),
//	})))
//}
//
//func CreateTable() {
//
//	_, err := dynamo.CreateTable(&dynamodb.CreateTableInput{
//		AttributeDefinitions: []*dynamodb.AttributeDefinition{
//			{
//				AttributeName: aws.String("Id"),
//				AttributeType: aws.String("N"),
//			},
//		},
//		KeySchema: []*dynamodb.KeySchemaElement{
//			{
//				AttributeName: aws.String("Id"),
//				KeyType:       aws.String("HASH"),
//			},
//		},
//		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
//			ReadCapacityUnits:  aws.Int64(1),
//			WriteCapacityUnits: aws.Int64(1),
//		},
//		TableName: aws.String(TABLE_NAME),
//	})
//
//	if err != nil {
//		if aerr, ok := err.(awserr.Error); ok {
//
//			fmt.Println(aerr.Error())
//		}
//	}
//}
//
//func PutItem(person Person) {
//
//	_, err := dynamo.PutItem(&dynamodb.PutItemInput{
//		Item: map[string]*dynamodb.AttributeValue{
//			"Id": {
//				N: aws.String(strconv.Itoa(person.Id)),
//			},
//			"Name": {
//				S: aws.String(person.Name),
//			},
//			"Website": {
//				S: aws.String(person.Website),
//			},
//		},
//		TableName: aws.String(TABLE_NAME),
//	})
//
//	if err != nil {
//		if aerr, ok := err.(awserr.Error); ok {
//
//			fmt.Println(aerr.Error())
//		}
//	}
//}
//
//func UpdateItem(person Person) {
//	_, err := dynamo.UpdateItem(&dynamodb.UpdateItemInput{
//		ExpressionAttributeNames: map[string]*string{
//			"#N": aws.String("Name"),
//			"#W": aws.String("Website"),
//		},
//		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
//			":Name": {
//				S: aws.String(person.Name),
//			},
//			":Website": {
//				S: aws.String(person.Website),
//			},
//		},
//		Key: map[string]*dynamodb.AttributeValue{
//			"Id": {
//				N: aws.String(strconv.Itoa(person.Id)),
//			},
//		},
//		TableName:        aws.String(TABLE_NAME),
//		UpdateExpression: aws.String("SET #N = :Name, #W = :Website"),
//	})
//
//	if err != nil {
//		if aerr, ok := err.(awserr.Error); ok {
//
//			fmt.Println(aerr.Error())
//		}
//	}
//}
//
//func DeleteItem(id int) {
//	_, err := dynamo.DeleteItem(&dynamodb.DeleteItemInput{
//		Key: map[string]*dynamodb.AttributeValue{
//			"Id": {
//				N: aws.String(strconv.Itoa(id)),
//			},
//		},
//		TableName: aws.String(TABLE_NAME),
//	})
//
//	if err != nil {
//		if aerr, ok := err.(awserr.Error); ok {
//
//			fmt.Println(aerr.Error())
//		}
//	}
//}
//
//func getItem(id int) (person Person) {
//	result, err := dynamo.GetItem(&dynamodb.GetItemInput{
//		Key: map[string]*dynamodb.AttributeValue{
//			"Id": {
//				N: aws.String(strconv.Itoa(id)),
//			},
//		},
//		TableName: aws.String(TABLE_NAME),
//	})
//
//	if err != nil {
//		if aerr, ok := err.(awserr.Error); ok {
//
//			fmt.Println(aerr.Error())
//		}
//	}
//
//	err = dynamodbattribute.UnmarshalMap(result.Item, &person)
//
//	if err != nil {
//		panic(err)
//	}
//
//	return person
//}
//
//func main() {
//
//	//e := echo.New()
//	//e.GET("/test", func(c echo.Context) error {
//	//	return c.String(http.StatusOK, "Hello, World!")
//	//})
//	//e.Logger.Fatal(e.Start(":1323"))
//
//	CreateTable()
//
//	var person Person = Person{
//		Id:      6,
//		Name:    "mina",
//		Website: "mina.ir",
//	}
//	PutItem(person)
//
//	fmt.Println(getItem(5))
//
//	var person2 Person = Person{
//		Id:      5,
//		Name:    "Majid 3",
//		Website: "majid.ir",
//	}
//
//	UpdateItem(person2)
//
//	DeleteItem(8)
//
//	fmt.Println(getItem(5))
//}
