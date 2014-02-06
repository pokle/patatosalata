package main

import (
  "fmt"
  "net/http"
  "os"
  "log"
  "io/ioutil"
)

func GET(url string) string {

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

  return string(result)

}

func main() {
  
  if len(os.Args) != 2 {
    log.Fatal("missing url")
  }
  
  url := os.Args[1]


  fmt.Printf("%s\n", GET(url))

}