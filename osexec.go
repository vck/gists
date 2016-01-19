package main

import (
  "fmt"
  "os/exec"
  "log"
)

func shell(cmd string)string{
  var (
    cmdOut []byte
    err error
    )

  if cmdOut, err = exec.Command(cmd).Output(); err != nil{
   log.Fatal(err)
  }
  return string(cmdOut)

}

func main() {
  str := shell("sasaasaa")
  fmt.Println(str)
}