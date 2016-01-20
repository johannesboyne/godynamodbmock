package godynamodbmock

import (
	"fmt"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type TestStruct struct {
	IdKey     string
	TestValue string
}

func TestNewMockCtrl(t *testing.T) {
	dynmodbservicemock := NewMockService(t)
	dynmodbservicemock.EXPECT().GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("testtable"),
		Key: map[string]*dynamodb.AttributeValue{
			"IdKey": &dynamodb.AttributeValue{
				S: aws.String("mykey"),
			},
		},
	}).Return(&dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"IdKey": &dynamodb.AttributeValue{
				S: aws.String("mykey"),
			},
			"TestValue": &dynamodb.AttributeValue{
				S: aws.String("test value"),
			},
		},
	}, nil)

	params := &dynamodb.GetItemInput{
		TableName: aws.String("testtable"),
		Key: map[string]*dynamodb.AttributeValue{
			"IdKey": &dynamodb.AttributeValue{
				S: aws.String("mykey"),
			},
		},
	}

	resp, err := dynmodbservicemock.GetItem(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	testObj := &TestStruct{}
	err = dynamodbattribute.ConvertFromMap(resp.Item, testObj)
	log.Printf("%+v\n", testObj)
}
