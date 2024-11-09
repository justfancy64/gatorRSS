package main

import(
  "fmt"
  "os"
  "database/sql"
  "github.com/justfancy64/gatorRSS/internal/config"
  "github.com/justfancy64/gatorRSS/internal/state"
  "github.com/justfancy64/gatorRSS/internal/commands"
  "github.com/justfancy64/gatorRSS/internal/database"


)
import _ "github.com/lib/pq"

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
  dbURL := st.Cfg.DbURL
  //fmt.Println(dbURL)

  db, err := sql.Open("postgres", dbURL)
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()

  dbQueries := database.New(db) // this is a pointer
  st.DB = dbQueries

  var cmds commands.Commands
  cmdmap := make(map[string]func(*state.State, commands.Command) error)
  cmds.CmdMap = cmdmap


  cmds.Register("login",    commands.HandlerLogins)
  cmds.Register("register", commands.RegisterHandler)
  cmds.Register("reset",    commands.HandlerClear)
  cmds.Register("users",    commands.HandlerListUsers)
  cmds.Register("agg",      commands.HandlerAgg)
  cmds.Register("addfeed",  commands.MiddleWareLoggedIn(commands.HandlerAddFeed))
  cmds.Register("feeds",    commands.HandlerListFeed)
  cmds.Register("follow",   commands.MiddleWareLoggedIn(commands.HandlerFollow))
  cmds.Register("following",commands.MiddleWareLoggedIn(commands.HandlerFollowing))
  cmds.Register("unfollow", commands.MiddleWareLoggedIn(commands.HandlerUnfollow))


  err = cmds.Run(&st, usercmd)



  if err != nil {

    fmt.Println(err)
  }


 
}
