-- +goose Up
CREATE TABLE auth.users (
  serial varchar PRIMARY KEY,
  username varchar,
  email varchar,
  name varchar,
  password varchar,
  role varchar,

  created_at timestamp,
  created_by varchar,
  updated_at timestamp,
  updated_by varchar,
  deleted_at timestamp,
  deleted_by varchar
);

CREATE UNIQUE INDEX idx_users_username_unique ON auth.users (username);
CREATE UNIQUE INDEX idx_users_email_unique ON auth.users (email);

-- +goose Down
DROP TABLE auth.users;