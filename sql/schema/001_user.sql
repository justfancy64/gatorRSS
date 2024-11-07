-- +goose Up
CREATE TABLE users(
  id         UUID      PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name       TEXT      NOT NULL UNIQUE
);

CREATE TABLE feeds(
  id           UUID           PRIMARY KEY,
  created_at   TIMESTAMP      NOT NULL,
  updated_at   TIMESTAMP      NOT NULL,
  name         TEXT           NOT NULL,
  url          TEXT           NOT NULL UNIQUE,
  user_id      UUID,         
  FOREIGN  KEY (user_id) 
  REFERENCES   users(id) ON DELETE CASCADE

);
-- +goose Down
DROP TABLE feeds;
DROP TABLE users;


