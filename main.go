package main

import (
  "fmt"
  "os"
  "os/user"
  "monke_lang_interpreter/repl"
)

func main() {
  user, err := user.Current()
  if err != nil {
    panic(err)
  }
  fmt.Printf("Hello %s! This is the MONKE PROG LANG!\n", user.Username)
  fmt.Printf("Feel free to type in some commands\n")
  repl.Start(os.Stdin, os.Stdout)
}

