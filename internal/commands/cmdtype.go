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
  handlerfunc, ok := c.CmdMap[cmd.Name] 
  if !ok {
    return fmt.Errorf("command not found")
  }
  return handlerfunc(s, cmd)
}
