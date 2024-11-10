
       Welcome to gatorRSS

a simple cli to tool for browsing subscribed rss feeds

Dependencies:
       golang 
       Postgres

Installation:
    create a file in your home named `.gatorConfig.json`
    copy paste
    `{
     "db_url":"postgres://postgres:Example",
     "current_user_name":"",
     "dbstring":"postgres://postgres:Example"
    }`
    where do_url is your dbs connection string
    clone this repo and run
    `$go install`
    


Usage:

       gatorRSS <command> [arguments]

The commands are:

       register  [user]      registers a new user
       reser                 resets the database[VERY DANGEROUS!]
       users                 lists all users registed
       update                updates feeds with new posts
       addfeed   [name,url]  adds new feed for the database
       follow    [url]       adds a feed to a users followed list
       following             lists current users followed feeds
       unfollow              unfollows feed for current user
       browse    [amount]    shows amount of posts
       read      [url]       shows description of post

