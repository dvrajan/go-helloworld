package main

import (
  "fmt"
  "net/http"
  "os"
)

const (
  port = ":80"
)

var calls = 0

func HelloWorld(w http.ResponseWriter, r *http.Request) {
  calls++
  fmt.Fprintf(w, "Hello, world! You have called me %d times (Config.ENV: %s).\n", calls,  os.Getenv("ENV"))
}

func init() {
  fmt.Printf("Started server at http://localhost%v.\n", port)
  http.HandleFunc("/", HelloWorld)
  http.ListenAndServe(port, nil)
}

func main() {}
