//go:generate mockgen -destination mock/dynamodbiface.go github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface DynamoDBAPI
package godynamodbmock

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/johannesboyne/godynamodbmock/mock"
)

func NewMockService(t *testing.T) *mock_dynamodbiface.MockDynamoDBAPI {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	serviceMock := mock_dynamodbiface.NewMockDynamoDBAPI(mockCtrl)
	return serviceMock
}
