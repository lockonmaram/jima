-- +goose Up
CREATE TABLE auth.user_groups (
  serial varchar PRIMARY KEY,
  user_serial varchar,
  group_serial varchar,
  user_group_role varchar,

  created_at timestamp,
  created_by varchar REFERENCES auth.users(serial),
  updated_at timestamp,
  updated_by varchar REFERENCES auth.users(serial),
  deleted_at timestamp,
  deleted_by varchar REFERENCES auth.users(serial)
);

ALTER TABLE auth.user_groups ADD FOREIGN KEY (user_serial) REFERENCES auth.users (serial);
ALTER TABLE auth.user_groups ADD FOREIGN KEY (group_serial) REFERENCES auth.groups (serial);

-- +goose Down
DROP TABLE auth.user_groups;
