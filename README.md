# Gosalesforce
Golang Salesforce REST API library. This library is inspired by the excellent "Simple-Salesforce" Python library.
Currently only OAuth2 is supported and an interface{} object is returned to keep things simple.

API documentation can be found at the [Godoc page](http://godoc.org/github.com/tux0010/gosalesforce)

## Examples

~~~ go
package main

import (
  "fmt"
  "encoding/json"
  "github.com/tux0010/gosalesforce"
)

func main() {
  instanceURL := "https://na23.salesforce.com"
  sessionID := "deadbeefbeefcafe"
  
  client := NewSalesforceClient(instanceURL, sessionID)
  data, err := client.Get("Someobject__C", "somerecordID")
  if err != nil {
    panic(err)
  }
  
  j, _ := json.MarshalIndent(data, "", " ")
  fmt.Println(string(j))
}
~~~
