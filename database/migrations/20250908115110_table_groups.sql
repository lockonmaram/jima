-- +goose Up
CREATE TABLE jima_auth.groups (
  serial varchar PRIMARY KEY,
  name varchar,

  created_at timestamp,
  created_by varchar,
  updated_at timestamp,
  updated_by varchar,
  deleted_at timestamp,
  deleted_by varchar
);

CREATE INDEX idx_groups_name ON jima_auth.groups (name);

-- +goose Down
DROP TABLE jima_auth.groups;