package main

import (
  "fmt"
  "github.com/0bvim/octobot/initialize" // dir name 
) 

func main() {
  token := initial.GetToken() // package name to call the functions
  fmt.Println("Success " + token)
}
