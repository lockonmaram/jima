-- +goose Up
CREATE TABLE auth.groups (
  serial varchar PRIMARY KEY,
  name varchar,
  created_at timestamp,
  updated_at timestamp,
  deleted_at timestamp
);

CREATE INDEX idx_groups_name ON auth.groups (name);

-- +goose Down
DROP TABLE auth.groups;