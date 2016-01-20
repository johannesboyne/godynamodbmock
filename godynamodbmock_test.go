package godynamodbmock

import (
	"fmt"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

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
	log.Println("test")

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

	// Pretty-print the response data.
	fmt.Println(resp)
}
