package main

import (
  "fmt"
  "net/http"
  "os"
  "log"
  "io/ioutil"
  "encoding/json"
)

func GET(url string) []byte {

  resp, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  } 

  defer resp.Body.Close()

  log.Printf("%s -- %s", url, resp)

  result, err := ioutil.ReadAll(resp.Body)
  
  if err != nil {
    log.Fatal(err)
  }

  return result
}

type Cluster struct {
  Name string
  Nodes int64
}

type Message struct {
  Clusters []Cluster
}


func parseMessage(bytes []byte) Message {
  var m Message
  err := json.Unmarshal(bytes, &m)
  if err != nil {
    log.Fatal(err)
  }
  return m
}

func main() {
  
  if len(os.Args) != 2 {
    log.Fatal("missing url")
  }
  
  url := os.Args[1]


  bytes := GET(url)
  // fmt.Printf("%s\n", bytes)
  fmt.Println(parseMessage(bytes))
}