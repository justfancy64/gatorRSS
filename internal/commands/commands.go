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
  fmt.Printf("user %s was created successfully\n",user.Name )
  //fmt.Println(user.ID)
  return nil
    
}


func HandlerClear(s *state.State, cmd Command) error {
  args := cmd.Args
  if len(args) > 0 {
    return fmt.Errorf("no arguments neededwith clear command")
  }
  err := s.DB.ClearPosts(context.Background())
  if err != nil {
    return fmt.Errorf("error clearing posts table: %v",err)
  }

 
  err = s.DB.ClearUser(context.Background())
  if err != nil {
    return fmt.Errorf("error clearing users table: %v",err)
  }
  err = s.Cfg.SetUser("")
  if err != nil {
    return err
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


func HandlerAgg(s *state.State,cmd Command) error {
  if len(cmd.Args) != 1 {
    return fmt.Errorf("agg commands needs a valid time duration eg: 1m 1s 1ms")
  }
  t , err := time.ParseDuration(cmd.Args[0])
  if err != nil {
    return fmt.Errorf("error in timeparseduration: %v", err)
  }

  ticker := time.NewTicker(t)

  for ; ; <-ticker.C {
    ScrapeFeeds(s)

  }
        

}
func HandlerAddFeed(s *state.State, cmd Command, user database.User) error {
  if len(cmd.Args) != 2 {
    return fmt.Errorf("not enough arguments needs: Name URL")
  }

   _, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
    ID:          uuid.New(),
    CreatedAt:   time.Now().UTC(),
    UpdatedAt:   time.Now().UTC(),
    Name:        cmd.Args[0],
    Url:         cmd.Args[1],
    UserID:      user.ID,


  })
  if err != nil {
    return fmt.Errorf("error in CreateFeedFunc: %v",err)
  }
  var newArgs  []string
  newArgs = append(newArgs, cmd.Args[1])
  newcmd := Command{
    Name:    cmd.Name,
    Args:    newArgs,
  }
  err = HandlerFollow(s, newcmd, user)
  if err != nil {
    return err
  }
  return nil
}



func HandlerListFeed(s *state.State, cmd Command) error{
  if len(cmd.Args) > 0 {
    return fmt.Errorf("no arguments needed with 'feeds' command")
   }
  feeds, err := s.DB.ListFeed(context.Background()) // []ListFeedRow{Name, Url,Name_2}
  if err != nil {
    return fmt.Errorf("error in ListFeed: %v", err)
  }
  fmt.Println(feeds)
  return nil
}



func HandlerFollow(s *state.State, cmd Command, user database.User) error{
  if len(cmd.Args) != 1 {
    return fmt.Errorf("not enough arguments need: URL")
  }

  feed, err := s.DB.GetFeed(context.Background(), cmd.Args[0])
  if err != nil {
    return fmt.Errorf("error fetching feed from db: %v",err)
  }


  row, err := s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
    ID:            uuid.New(),
    CreatedAt:     time.Now().UTC(),
    UpdatedAt:     time.Now().UTC(),
    UserID:        user.ID,
    FeedID:        feed.ID,
  }) 
  if err != nil {
    return fmt.Errorf("error in CreateFeedFollow: %v", err)
  }
  fmt.Printf("Feed: %s added for user: %s\n",row.FeedName, user.Name)
  return nil

}


func HandlerFollowing(s *state.State, cmd Command, user database.User) error{
  if len(cmd.Args) != 0 {
    return fmt.Errorf("Following commands takes no arguments")
  }
  feeds, err := s.DB.GetUserFollows(context.Background(), user.ID)
  if err != nil {
    return fmt.Errorf("error in GetUserFollows: %v", err)
  }
  for _,feed := range feeds {
    fmt.Println(feed)
  }
  return nil
}

func HandlerUnfollow(s *state.State, cmd Command, user database.User) error {
  if len(cmd.Args) != 1 {
    return fmt.Errorf("provide url of feed u wish to unfollow")
  }
  feed, err := s.DB.GetFeed(context.Background(), cmd.Args[0])
  if err != nil {
    return err
  }
  err = s.DB.DeleteFollow(context.Background(), database.DeleteFollowParams{
     UserID:      user.ID,
     FeedID:      feed.ID,
  })
  if err != nil {
    return err
  }

  return nil
}

// middleware loggin checker for handlers requiring user to be logged in

func MiddleWareLoggedIn(handler func(s *state.State, cmd Command, user database.User) error) func(*state.State, Command) error {
  return func(s *state.State, cmd Command) error {
  CurrUser, err := s.DB.GetUser(context.Background(), s.Cfg.CurrentUserName)
  if err != nil {
    return  fmt.Errorf("error fetching user info from DB: %v",err)
  }
    return handler(s, cmd, CurrUser)
  }
}


//db scrape helper function
func ScrapeFeeds(s *state.State) {
  url, err := s.DB.GetNextFeedToFetch(context.Background())
  if err != nil {
     fmt.Errorf("error in GetNextFeedToFetch: %v", err)
  }
  dbfeed, err := s.DB.GetFeed(context.Background(),url)
  if err != nil {
     fmt.Errorf("error fetching feed: %v",err)
  }
  err = s.DB.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
    UpdatedAt:       time.Now().UTC(),
    ID:              dbfeed.ID,

})
  if err != nil {
     fmt.Errorf("error in marking feed as fetched: %v", err)
  }

  rssfeed, err := rss.FetchFeed(context.Background(),url)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Printf("adding post from %s to db\n", rssfeed.Channel.Link)
  timepub, err := time.Parse("2024-11-09 11:47:42.00913", rssfeed.Channel.Item[0].PubDate)

  if err != nil {
    fmt.Errorf("error in time.parse: %v", err)
  }

  for _,item := range rssfeed.Channel.Item {

  _,err := s.DB.CreatePost(context.Background(), database.CreatePostParams{
    ID:             uuid.New(),
    CreatedAt:      time.Now().UTC(),
    UpdatedAt:      time.Now().UTC(),
    Title:          item.Title,
    Url:            item.Link,
    Description:    item.Description,
    PublishedAt:    timepub,
    FeedID:         dbfeed.ID,
  })
  if err != nil {
     fmt.Errorf("error in CreatePost quary")
  }

  }
}
