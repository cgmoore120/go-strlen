package main

import(
      "fmt"
      "unicode/utf8"
      "encoding/json"
      "os"
)

func main() {
  byte_map := map[string]int{"length" : utf8.RuneCountInString(os.Args[1])}

  result, _ := json.Marshal(byte_map)
  fmt.Println(string(result))
}
