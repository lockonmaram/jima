-- +goose Up
CREATE TABLE jima_auth.user_groups (
  serial varchar PRIMARY KEY,
  user_serial varchar,
  group_serial varchar,
  role varchar,

  created_at timestamp,
  created_by varchar,
  updated_at timestamp,
  updated_by varchar,
  deleted_at timestamp,
  deleted_by varchar
);

ALTER TABLE jima_auth.user_groups ADD FOREIGN KEY (user_serial) REFERENCES jima_auth.users (serial);
ALTER TABLE jima_auth.user_groups ADD FOREIGN KEY (group_serial) REFERENCES jima_auth.groups (serial);

-- +goose Down
DROP TABLE jima_auth.user_groups;
