package commands

import (
  "fmt"
  "context"
  "time"
  "github.com/justfancy64/gatorRSS/internal/state"
  "github.com/justfancy64/gatorRSS/internal/database"
  "github.com/justfancy64/gatorRSS/internal/rss"
  "github.com/google/uuid"
)



func HandlerLogins(s *state.State, cmd Command) error {
  if len(cmd.Args) == 0 {
    return fmt.Errorf("no username given")
  }

  _, err := s.DB.GetUser(context.Background(), cmd.Args[0])
    if err != nil {
  return fmt.Errorf("usage: %s <name>", cmd.Name)
  }
  err = s.Cfg.SetUser(cmd.Args[0])
  if err != nil {
    return err
  }
  fmt.Println("user has been set")
  return nil
}

  
  
func RegisterHandler(s *state.State, cmd Command) error {
  if len(cmd.Args) < 1 {
    return fmt.Errorf("No name was passed in registration")
  }

  id := uuid.New()

  fmt.Println("adding user to db")
  t := time.Now()

  user, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
  ID:        id,
  CreatedAt: t,
  UpdatedAt: t,
  Name:      cmd.Args[0],
  })
  if err != nil {
    fmt.Println(err)
    return fmt.Errorf("error adding user to database")

  }


  err = s.Cfg.SetUser(cmd.Args[0])
  if err != nil {
    return err
  }
  fmt.Printf("user %s was created successfully",user.Name )
  //fmt.Println(user.ID)
  return nil
    
}


func HandlerClear(s *state.State, cmd Command) error {
  args := cmd.Args
  if len(args) > 0 {
    return fmt.Errorf("no arguments neededwith clear command")
  }

 
  err := s.DB.ClearUser(context.Background())
  if err != nil {
    return fmt.Errorf("error clearing users table: %v",err)
  }
  return nil

}

func HandlerListUsers(s *state.State, cmd Command) error {
  args := cmd.Args
  if len(args) > 0 {
    return fmt.Errorf("no arguments neededwith clear command")
  }
  userlist, err := s.DB.ListUsers(context.Background())
  if err != nil {
    return fmt.Errorf("error clearing users table: %v",err)
  }
  for _,user := range userlist {
    if user == s.Cfg.CurrentUserName {
      user = user + " (current)"
    }
    fmt.Println(user)
  }

  return nil


}


func HandlerAgg(s *state.State, cmd Command) error {
  // rss feed testing
  ctx := context.Background()
  feed, err := rss.FetchFeed(ctx, "https://www.wagslane.dev/index.xml")
  if err != nil {
    return err
  }
  fmt.Println(feed)
  return nil
}

func HandlerAddFeed(s *state.State, cmd Command) error {
  if len(cmd.Args) < 2 {
    return fmt.Errorf("not enough arguments needs: Name URL")
  }
  
  CurrUser, err := s.DB.GetUser(context.Background(), s.Cfg.CurrentUserName)
  if err != nil {
    fmt.Errorf("error fetching user info from DB: %v",err)
  }
  

   feed, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
    ID:          uuid.New(),
    CreatedAt:   time.Now().UTC(),
    UpdatedAt:   time.Now().UTC(),
    Name:        cmd.Args[0],
    Url:         cmd.Args[1],
    UserID:       CurrUser.ID,


  })
  if err != nil {
    return fmt.Errorf("error in CreateFeedFunc: %v",err)
  }
  fmt.Println(feed)
  return nil
}



func HandlerListFeed(s *state.State, cmd Command) error{
  if len(cmd.Args) > 0 {
    return fmt.Errorf("no arguments needed with 'feeds' command")
   }
  feeds, err := s.DB.ListFeed(context.Background()) // []ListFeedRow{Name, Url,Name_2}
  if err != nil {
    return fmt.Errorf("error in ListFeed: %v")
  }
  fmt.Println(feeds)
  return nil
}
