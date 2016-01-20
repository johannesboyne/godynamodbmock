// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface (interfaces: DynamoDBAPI)

package mock_dynamodbiface

import (
	request "github.com/aws/aws-sdk-go/aws/request"
	dynamodb "github.com/aws/aws-sdk-go/service/dynamodb"
	gomock "github.com/golang/mock/gomock"
)

// Mock of DynamoDBAPI interface
type MockDynamoDBAPI struct {
	ctrl     *gomock.Controller
	recorder *_MockDynamoDBAPIRecorder
}

// Recorder for MockDynamoDBAPI (not exported)
type _MockDynamoDBAPIRecorder struct {
	mock *MockDynamoDBAPI
}

func NewMockDynamoDBAPI(ctrl *gomock.Controller) *MockDynamoDBAPI {
	mock := &MockDynamoDBAPI{ctrl: ctrl}
	mock.recorder = &_MockDynamoDBAPIRecorder{mock}
	return mock
}

func (_m *MockDynamoDBAPI) EXPECT() *_MockDynamoDBAPIRecorder {
	return _m.recorder
}

func (_m *MockDynamoDBAPI) BatchGetItem(_param0 *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	ret := _m.ctrl.Call(_m, "BatchGetItem", _param0)
	ret0, _ := ret[0].(*dynamodb.BatchGetItemOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) BatchGetItem(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BatchGetItem", arg0)
}

func (_m *MockDynamoDBAPI) BatchGetItemPages(_param0 *dynamodb.BatchGetItemInput, _param1 func(*dynamodb.BatchGetItemOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "BatchGetItemPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDynamoDBAPIRecorder) BatchGetItemPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BatchGetItemPages", arg0, arg1)
}

func (_m *MockDynamoDBAPI) BatchGetItemRequest(_param0 *dynamodb.BatchGetItemInput) (*request.Request, *dynamodb.BatchGetItemOutput) {
	ret := _m.ctrl.Call(_m, "BatchGetItemRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dynamodb.BatchGetItemOutput)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) BatchGetItemRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BatchGetItemRequest", arg0)
}

func (_m *MockDynamoDBAPI) BatchWriteItem(_param0 *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
	ret := _m.ctrl.Call(_m, "BatchWriteItem", _param0)
	ret0, _ := ret[0].(*dynamodb.BatchWriteItemOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) BatchWriteItem(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BatchWriteItem", arg0)
}

func (_m *MockDynamoDBAPI) BatchWriteItemRequest(_param0 *dynamodb.BatchWriteItemInput) (*request.Request, *dynamodb.BatchWriteItemOutput) {
	ret := _m.ctrl.Call(_m, "BatchWriteItemRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dynamodb.BatchWriteItemOutput)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) BatchWriteItemRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BatchWriteItemRequest", arg0)
}

func (_m *MockDynamoDBAPI) CreateTable(_param0 *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
	ret := _m.ctrl.Call(_m, "CreateTable", _param0)
	ret0, _ := ret[0].(*dynamodb.CreateTableOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) CreateTable(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateTable", arg0)
}

func (_m *MockDynamoDBAPI) CreateTableRequest(_param0 *dynamodb.CreateTableInput) (*request.Request, *dynamodb.CreateTableOutput) {
	ret := _m.ctrl.Call(_m, "CreateTableRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dynamodb.CreateTableOutput)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) CreateTableRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateTableRequest", arg0)
}

func (_m *MockDynamoDBAPI) DeleteItem(_param0 *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	ret := _m.ctrl.Call(_m, "DeleteItem", _param0)
	ret0, _ := ret[0].(*dynamodb.DeleteItemOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) DeleteItem(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteItem", arg0)
}

func (_m *MockDynamoDBAPI) DeleteItemRequest(_param0 *dynamodb.DeleteItemInput) (*request.Request, *dynamodb.DeleteItemOutput) {
	ret := _m.ctrl.Call(_m, "DeleteItemRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dynamodb.DeleteItemOutput)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) DeleteItemRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteItemRequest", arg0)
}

func (_m *MockDynamoDBAPI) DeleteTable(_param0 *dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error) {
	ret := _m.ctrl.Call(_m, "DeleteTable", _param0)
	ret0, _ := ret[0].(*dynamodb.DeleteTableOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) DeleteTable(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteTable", arg0)
}

func (_m *MockDynamoDBAPI) DeleteTableRequest(_param0 *dynamodb.DeleteTableInput) (*request.Request, *dynamodb.DeleteTableOutput) {
	ret := _m.ctrl.Call(_m, "DeleteTableRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dynamodb.DeleteTableOutput)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) DeleteTableRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteTableRequest", arg0)
}

func (_m *MockDynamoDBAPI) DescribeTable(_param0 *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error) {
	ret := _m.ctrl.Call(_m, "DescribeTable", _param0)
	ret0, _ := ret[0].(*dynamodb.DescribeTableOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) DescribeTable(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeTable", arg0)
}

func (_m *MockDynamoDBAPI) DescribeTableRequest(_param0 *dynamodb.DescribeTableInput) (*request.Request, *dynamodb.DescribeTableOutput) {
	ret := _m.ctrl.Call(_m, "DescribeTableRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dynamodb.DescribeTableOutput)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) DescribeTableRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeTableRequest", arg0)
}

func (_m *MockDynamoDBAPI) GetItem(_param0 *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	ret := _m.ctrl.Call(_m, "GetItem", _param0)
	ret0, _ := ret[0].(*dynamodb.GetItemOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) GetItem(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetItem", arg0)
}

func (_m *MockDynamoDBAPI) GetItemRequest(_param0 *dynamodb.GetItemInput) (*request.Request, *dynamodb.GetItemOutput) {
	ret := _m.ctrl.Call(_m, "GetItemRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dynamodb.GetItemOutput)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) GetItemRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetItemRequest", arg0)
}

func (_m *MockDynamoDBAPI) ListTables(_param0 *dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error) {
	ret := _m.ctrl.Call(_m, "ListTables", _param0)
	ret0, _ := ret[0].(*dynamodb.ListTablesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) ListTables(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTables", arg0)
}

func (_m *MockDynamoDBAPI) ListTablesPages(_param0 *dynamodb.ListTablesInput, _param1 func(*dynamodb.ListTablesOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "ListTablesPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDynamoDBAPIRecorder) ListTablesPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTablesPages", arg0, arg1)
}

func (_m *MockDynamoDBAPI) ListTablesRequest(_param0 *dynamodb.ListTablesInput) (*request.Request, *dynamodb.ListTablesOutput) {
	ret := _m.ctrl.Call(_m, "ListTablesRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dynamodb.ListTablesOutput)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) ListTablesRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTablesRequest", arg0)
}

func (_m *MockDynamoDBAPI) PutItem(_param0 *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	ret := _m.ctrl.Call(_m, "PutItem", _param0)
	ret0, _ := ret[0].(*dynamodb.PutItemOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) PutItem(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutItem", arg0)
}

func (_m *MockDynamoDBAPI) PutItemRequest(_param0 *dynamodb.PutItemInput) (*request.Request, *dynamodb.PutItemOutput) {
	ret := _m.ctrl.Call(_m, "PutItemRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dynamodb.PutItemOutput)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) PutItemRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutItemRequest", arg0)
}

func (_m *MockDynamoDBAPI) Query(_param0 *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	ret := _m.ctrl.Call(_m, "Query", _param0)
	ret0, _ := ret[0].(*dynamodb.QueryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) Query(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Query", arg0)
}

func (_m *MockDynamoDBAPI) QueryPages(_param0 *dynamodb.QueryInput, _param1 func(*dynamodb.QueryOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "QueryPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDynamoDBAPIRecorder) QueryPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "QueryPages", arg0, arg1)
}

func (_m *MockDynamoDBAPI) QueryRequest(_param0 *dynamodb.QueryInput) (*request.Request, *dynamodb.QueryOutput) {
	ret := _m.ctrl.Call(_m, "QueryRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dynamodb.QueryOutput)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) QueryRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "QueryRequest", arg0)
}

func (_m *MockDynamoDBAPI) Scan(_param0 *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	ret := _m.ctrl.Call(_m, "Scan", _param0)
	ret0, _ := ret[0].(*dynamodb.ScanOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) Scan(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Scan", arg0)
}

func (_m *MockDynamoDBAPI) ScanPages(_param0 *dynamodb.ScanInput, _param1 func(*dynamodb.ScanOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "ScanPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDynamoDBAPIRecorder) ScanPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ScanPages", arg0, arg1)
}

func (_m *MockDynamoDBAPI) ScanRequest(_param0 *dynamodb.ScanInput) (*request.Request, *dynamodb.ScanOutput) {
	ret := _m.ctrl.Call(_m, "ScanRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dynamodb.ScanOutput)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) ScanRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ScanRequest", arg0)
}

func (_m *MockDynamoDBAPI) UpdateItem(_param0 *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	ret := _m.ctrl.Call(_m, "UpdateItem", _param0)
	ret0, _ := ret[0].(*dynamodb.UpdateItemOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) UpdateItem(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateItem", arg0)
}

func (_m *MockDynamoDBAPI) UpdateItemRequest(_param0 *dynamodb.UpdateItemInput) (*request.Request, *dynamodb.UpdateItemOutput) {
	ret := _m.ctrl.Call(_m, "UpdateItemRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dynamodb.UpdateItemOutput)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) UpdateItemRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateItemRequest", arg0)
}

func (_m *MockDynamoDBAPI) UpdateTable(_param0 *dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error) {
	ret := _m.ctrl.Call(_m, "UpdateTable", _param0)
	ret0, _ := ret[0].(*dynamodb.UpdateTableOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) UpdateTable(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateTable", arg0)
}

func (_m *MockDynamoDBAPI) UpdateTableRequest(_param0 *dynamodb.UpdateTableInput) (*request.Request, *dynamodb.UpdateTableOutput) {
	ret := _m.ctrl.Call(_m, "UpdateTableRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dynamodb.UpdateTableOutput)
	return ret0, ret1
}

func (_mr *_MockDynamoDBAPIRecorder) UpdateTableRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateTableRequest", arg0)
}
