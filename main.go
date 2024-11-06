package main

import(
  "fmt"
  "os"
  "github.com/justfancy64/gatorRSS/internal/config"
  "github.com/justfancy64/gatorRSS/internal/state"
  "github.com/justfancy64/gatorRSS/internal/commands"


)


func main() {
  if len(os.Args) < 2 {
    fmt.Println("Error less than 1 argument given")
    os.Exit(1)
    return
  }


  var usercmd commands.Command
  usercmd.Name = os.Args[1]
  usercmd.Args = os.Args[2:]

  var st state.State

  cfg, err := config.Read()
  if err != nil {
    fmt.Println(err)
  }
  st.Cfg = &cfg

  var cmds commands.Commands
  cmdmap := make(map[string]func(*state.State, commands.Command) error)
  cmds.CmdMap = cmdmap
  cmds.Register(os.Args[1],  commands.HandlerLogins)
  err = cmds.Run(&st, usercmd)



  st.Cfg = &cfg
  if err != nil {

    fmt.Println(err)
    os.Exit(1)
  }


  return
 
}
