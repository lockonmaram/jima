-- +goose Up
ALTER TABLE jima_auth.users ADD COLUMN password_token varchar;
CREATE UNIQUE INDEX idx_users_password_token_unique ON jima_auth.users (password_token);

-- +goose Down
DROP INDEX jima_auth.idx_users_password_token_unique;
ALTER TABLE jima_auth.users DROP COLUMN password_token;
