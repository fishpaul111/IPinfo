package main

import (
  "encoding/json"
  "fmt"
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
  "io/ioutil"
  "net"
  "net/http"
  "strings"
)

// ****************************************************************************
// CONSTANTS
// ****************************************************************************
const URL = "https://json.geoiplookup.io/"
const PREFIX_TO_TRIM = "main.IPInfo"
const CHARS_TO_TRIM = "{}"


// ****************************************************************************
// TYPE DEFINES
// ****************************************************************************
type IPInfo struct {
  IP string `json:"ip"`
  ISP string `json:"isp"`
  HostName string `json:"hostname"`
  Latitude float64 `json:"latitude"`
  Longitude float64 `json:"longitude"`
  PostalCode string `json:"postal_code"`
  City string `json:"city"`
  CountryName string `json:"country_name"`
  ConnectionType string `json:"connection_type"`
  Success bool `json:"success"`
}


// ****************************************************************************
// UTILITIES
// ****************************************************************************
func GetIPInfo(ip string) IPInfo {
  var ipData IPInfo

  if nil != net.ParseIP(ip) {
    httpResponse, error := http.Get(URL + ip)
    if error != nil {
      fmt.Print(error.Error())
    }

    data, error := ioutil.ReadAll(httpResponse.Body)
    if error != nil {
      fmt.Print(error.Error())
    }
    defer httpResponse.Body.Close()

    error = json.Unmarshal(data, &ipData)
    if error != nil {
      fmt.Print(error.Error())
    }
  }

  return ipData
}


// ****************************************************************************
// HANDLER
// ****************************************************************************
func Engine(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  var response events.APIGatewayProxyResponse
  ip, success := request.QueryStringParameters["ip"]

  if success {
    response.StatusCode = http.StatusOK
    response.Headers = map[string]string{"Content-Type": "text/html; charset=utf-8"}
    response.Body = strings.Trim(strings.TrimPrefix(fmt.Sprintf("%#v", GetIPInfo(ip)), PREFIX_TO_TRIM), CHARS_TO_TRIM)
  } else {
    response.StatusCode = http.StatusBadRequest
  }

  return response, nil
}

// ****************************************************************************
// MAIN
// ****************************************************************************
func main() {
  lambda.Start(Engine)
}
