package main

import(
  "fmt"
  "github.com/justfancy64/gatorRSS/internal/config"
)


func main() {
  cfg, err := config.Read()
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println("writing username to config")
  err = cfg.SetUser("zach")
  if err != nil {
    fmt.Println(err)
  }

  config2, err := config.Read()
  fmt.Println(config2.DbURL)
  fmt.Println(config2.CurrentUserName)

  return
 
}
