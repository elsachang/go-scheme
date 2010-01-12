package main

import (
  fmt "fmt"
  // io "io"
)


type ObjectType int
const (
  _ = iota
  FIXNUM
)

type Object struct {
  tipe ObjectType
}

func main() {
  fmt.Printf("Welcome to Bootstrap Scheme. ")
  fmt.Printf("Use ctrl-c to exit.\n")

  for {

  }
}
