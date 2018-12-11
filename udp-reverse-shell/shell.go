package main

import (
  "fmt"
  "os/exec"
)

func CommandExec(cmd string) string {
  output, err := exec.Command("bash", "-c", cmd).Output()
  if err != nil {
    fmt.Println(err)
  }
  return string(output)
}

