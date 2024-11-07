package config

import (
  "fmt"
  "os"
  "encoding/json"
)

const configFIleName = "/.gatorconfig.json"


type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
	Dbstring        string `json:"dbstring"`
}




func Read() (Config, error) {

  filePath, err := getConfigFilePath()
  if err != nil {
    return Config{},fmt.Errorf("error getting config filepath: %v", err)

  }
  
  data, err := os.ReadFile(filePath)
  if err != nil {
    return Config{}, fmt.Errorf("error reading config: %v", err)
  }
  var cfg Config
  err = json.Unmarshal(data, &cfg)
  if err != nil {
    return Config{}, fmt.Errorf("error unmashaling data: %v", err)
  }

  return cfg, nil
}

func (c Config) SetUser(username string) error {

  filePath, err :=  getConfigFilePath()
  if err != nil {
    return fmt.Errorf("error getting config filepath: %v", err)
  }

  file, err := os.Create(filePath)
  if err != nil {
    return fmt.Errorf("error creating *file in SetUser: %v", err)
  }

  defer file.Close()


  c.CurrentUserName = username
  data, err := json.Marshal(c)
  if err != nil {
    return fmt.Errorf("error in SetUser: %v", err)
  }
  _, err = file.WriteString(string(data))
  if err != nil {
    return fmt.Errorf("error in SetUser while writing configfile: %v", err)
  }
  return nil

}



func getConfigFilePath() (string, error) {
  home, err := os.UserHomeDir()
  if err != nil {
    return  "", fmt.Errorf("error reading configfile: %v", err)
  }
  return home + configFIleName, nil
}

