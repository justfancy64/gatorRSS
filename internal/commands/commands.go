package commands

import (
  "fmt"
  "github.com/justfancy64/gatorRSS/internal/state"
)



func HandlerLogins(s *state.State, cmd Command) error {
  if len(cmd.Args) == 0 {
    return fmt.Errorf("no username given")
  }
  err := s.Cfg.SetUser(cmd.Args[0])
  if err != nil {
    return err
  }
  fmt.Println("user has been set")
  return nil
}
