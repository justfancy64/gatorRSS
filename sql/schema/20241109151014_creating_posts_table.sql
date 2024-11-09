-- +goose Up
-- +goose StatementBegin
create table posts(
  id       UUID   PRIMARY key,
  created_at    TIMESTAMP not null,
  updated_at    TIMESTAMP not null,
  title         text      not null,
  url           text      not null,
  description   text      not null,
  published_at  TIMESTAMP not null,
  feed_id       UUID      not null,
  foreign key   (feed_id)
  references    feeds(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table posts;
-- +goose StatementEnd
