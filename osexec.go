package main

import (
  "fmt"
  "os"
  "os/exec"
)

func shell(cmd string)string{
  var (
    cmdOut []byte
    err error
    )

  if cmdOut, err = exec.Command(cmd).Output(); err != nil{
    os.Exit(1)
  }
  return string(cmdOut)

}

func main() {
  str := shell("ls")
  fmt.Println(str)
}