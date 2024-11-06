package commands


import (
  "fmt"
  "github.com/justfancy64/gatorRSS/internal/state"
)



type Command struct {
  Name       string
  Args       []string
}


type Commands struct {
  CmdMap        map[string]func(*state.State, Command) error
}





func (c *Commands) Register(name string, f func(*state.State, Command) error) {
  c.CmdMap[name] = f

}


func (c *Commands) Run(s *state.State, cmd Command) error {
  handlerfunc := c.CmdMap[cmd.Name] 
  err := handlerfunc(s, cmd)
  if err != nil {
    return fmt.Errorf("error in %s command: %v",cmd.Name, err)
  }
  return nil
}
