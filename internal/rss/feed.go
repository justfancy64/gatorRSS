package rss

import (
  "fmt"
  "net/http"
  "encoding/xml"
  "io"
  "context"
  "html"

  )




func FetchFeed(ctx context.Context, feedURL string)  (*RSSFeed, error) {
  req, err := http.NewRequestWithContext(ctx,"GET",feedURL, nil)
  if err != nil {
  return nil, fmt.Errorf("error in FetchFeed on NewRequestWithContext: %v",err)
  }
  req.Header.Add("User-Agent","gator")

  client := &http.Client{}
  res, err := client.Do(req)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()

  body, err := io.ReadAll(res.Body)
  if err != nil {
    return nil, err
  }
 // fmt.Println(string(body))
  var rss RSSFeed
  err = xml.Unmarshal(body, &rss)
  if err != nil {
    return nil, err
  }

  htmlCleanup(&rss)
  return &rss, nil 
  
}



func htmlCleanup(f *RSSFeed){
  f.Channel.Title = html.UnescapeString(f.Channel.Title)
  f.Channel.Description = html.UnescapeString(f.Channel.Description)
  for _,item := range f.Channel.Item {
    item.Title = html.UnescapeString(item.Title)
    item.Description = html.UnescapeString(item.Description)
  }
}
