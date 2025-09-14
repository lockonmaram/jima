-- +goose Up
CREATE TABLE auth.groups (
  serial varchar PRIMARY KEY,
  name varchar,

  created_at timestamp,
  created_by varchar REFERENCES auth.users(serial),
  updated_at timestamp,
  updated_by varchar REFERENCES auth.users(serial),
  deleted_at timestamp,
  deleted_by varchar REFERENCES auth.users(serial)
);

CREATE INDEX idx_groups_name ON auth.groups (name);

-- +goose Down
DROP TABLE auth.groups;