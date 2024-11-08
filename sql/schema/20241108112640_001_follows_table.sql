-- +goose Up
-- +goose StatementBegin
create table feed_follows (
  id          UUID      primary key,
  created_at  TIMESTAMP NOT NULL,
  updated_at  TIMESTAMP NOT NULL,
  user_id     UUID      NOT NULL,
  feed_id     UUID      NOT NULL,
  UNIQUE      (user_id, feed_id),
  FOREIGN KEY (user_id)
  REFERENCES  users(id) ON DELETE CASCADE,
  FOREIGN KEY (feed_id)
  REFERENCES  feeds(id)  ON DELETE CASCADE
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
drop table feed_follows;
-- +goose StatementEnd
