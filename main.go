package main

import (
  "context"
  "fmt"
  "os"

  "github.com/ohhfishal/notify/cmd"
)

func main() {
  ctx := context.Background()
  if err := cmd.Run(ctx, os.Stdin, os.Stdout, os.Args[1:]); err != nil {
    fmt.Println(os.Stderr, err.Error())
    os.Exit(1)
  }
}
