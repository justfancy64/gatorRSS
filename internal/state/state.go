package state

import(
 "github.com/justfancy64/gatorRSS/internal/config"
 "github.com/justfancy64/gatorRSS/internal/database"
)



type State struct {
  DB   *database.Queries
  Cfg  *config.Config
}
