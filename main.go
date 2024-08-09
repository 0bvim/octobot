package main

import (
  "fmt"
  "github.com/0bvim/octobot/initialize" // dir name 
) 

func main() {
  token, err := initial.GetToken() // package name to call the functions
  if err != nil {
    return
  }
  fmt.Println("Success " + token)
}
