package main

import (
  "github.com/aws/aws-lambda-go/events"
  "testing"
  )


func Test_GetIPInfo_Base(test *testing.T) {
  targetResult := "Cloudflare, Inc."
  testResult := GetIPInfo("1.1.1.1")

  if testResult.ISP != targetResult {
  	test.Error("TEST FAILED. GetIPInfo returned the incorrect information.")
  }
}

func Test_GetIPInfo_Empty(test *testing.T) {
  targetResult := ""
  testResult := GetIPInfo("")

  if testResult.ISP != targetResult {
  	test.Error("TEST FAILED. GetIPInfo returned the incorrect information.")
  }
}

func Test_Engine(test *testing.T) {
  var testRequest events.APIGatewayProxyRequest
  var testResponse events.APIGatewayProxyResponse

  testResponse.StatusCode = 123
  testResponse, error := Engine(testRequest)

  if testResponse.StatusCode == 123 {
  	test.Error("TEST FAILED. Engine did not return a response.")
  }
  if error != nil {
    test.Error("TEST FAILED. Engine returned an error.")
  }
}
