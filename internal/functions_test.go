package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
    "github.com/gin-gonic/gin"
)

type mockDoErrorHTTPClient struct {}
type mockEmptyBodyReqHTTPClient struct {}
type mockInvalidTokenInBodyHTTPClient struct {}
type mockSuccessClient struct {}

func setupMockResponse(status string, statusCode int, bodyData interface{}) (*http.Response, error) {
    var body io.ReadCloser

    dataBytes, err := json.Marshal(bodyData)
    if err != nil {
        return nil, err
    }

    body = ioutil.NopCloser(bytes.NewBuffer(dataBytes))

    res := &http.Response{
        Status: status,
        StatusCode: statusCode,
        Body: body,
    }

    return res, nil
}

func (client *mockDoErrorHTTPClient) Do(req *http.Request) (*http.Response, error) {
    return nil, errors.New("Error")
}

func (client *mockEmptyBodyReqHTTPClient) Do(req *http.Request) (*http.Response, error) {
    res, err := setupMockResponse("200 OK", http.StatusOK, "")

    return res, err
}

func (client *mockInvalidTokenInBodyHTTPClient) Do(req *http.Request) (*http.Response, error) {
    bodyData := &SnipcartValidation{
        Token: "testToken",
    }

    res, err := setupMockResponse("200 OK", http.StatusOK, bodyData)

    return res, err
}

func (client *mockSuccessClient) Do(req *http.Request) (*http.Response, error) {
    bodyData := &SnipcartValidation{
        Token: "testValidToken",
    }

    res, err := setupMockResponse("200 OK", http.StatusOK, bodyData)

    return res, err
}


func TestCheckIfIntIsZeroValue(t *testing.T) {
    assert.False(t, checkIfIntIsZeroValue(1))
    assert.True(t, checkIfIntIsZeroValue(0))

    var i int

    assert.True(t, checkIfIntIsZeroValue(i))
}

func TestCheckIfStringIsZeroValue(t *testing.T) {
    assert.False(t, checkIfStringIsZeroValue("str"))
    assert.True(t, checkIfStringIsZeroValue(""))

    var str string

    assert.True(t, checkIfStringIsZeroValue(str))
}

func TestAppendToStringArray(t *testing.T) {
    var arr []string

    appendToStringArray("test", &arr)

    assert.True(t, len(arr) == 1)
}

func TestValidateSnipcartTokenDoError(t *testing.T) {
    client := &mockDoErrorHTTPClient{}
    token := "test"
    expectedError := "Error"
    err := validateSnipcartToken(client, token)

    assert.Equal(t, expectedError, err.Error())
}

func TestValidateSnipcartTokenEmptyBodyError(t *testing.T) {
    client := &mockEmptyBodyReqHTTPClient{}
    token := "test"
    expectedError := "json: cannot unmarshal string into Go value of type main.SnipcartValidation"
    err := validateSnipcartToken(client, token)

    assert.Equal(t, expectedError, err.Error())
}

func TestValidateSnipcartTokenErr2(t *testing.T) {
    client := &mockInvalidTokenInBodyHTTPClient{}
    token := "notTestToken"
    expectedError := "functions.go:64 - Token data compromised."
    err := validateSnipcartToken(client, token)

    assert.Equal(t, expectedError, err.Error())
}

func TestValidateSnipcartTokenSuccess(t *testing.T) {

    client := &mockInvalidTokenInBodyHTTPClient{}
    token := "testToken"
    err := validateSnipcartToken(client, token)

    assert.Equal(t, nil, err)
}

func setupServer(client Doer) *Server {
    // Will get initiallized to zero values
    var emailLogger *EmailLogger
    var router *gin.Engine
    var limiter *IPRateLimiter

    var server *Server = createServer(client, emailLogger, router, limiter)

    return server
}

func TestCallItemFunctionNoMetadata(t *testing.T) {
    mockClient := &mockSuccessClient{}

    server := setupServer(mockClient)

    metadata := &Metadata{}

    testItem := &Item{
        Metadata: *metadata,
    }

    sJSON := &SnipcartData{}

    err := server.callItemFunction(testItem, sJSON)
    assert.Nil(t, err)
}

func TestCallItemFunctionHandlePymail(t *testing.T) {
    mockClient := mockSuccessClient{}

    server := setupServer(mockClient)

    sendto := []string{"1", "test"}

    metadata := &Metadata{
        Sendto: sendto,
    }

    testItem := &Item{
        Metadata: *metadata,
    }

    sJSON := &SnipcartData{}

    err := server.callItemFunction(testItem, sJSON)
    fmt.Printf("Error is:%v and type:%T", err,err)
}
