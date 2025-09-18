-- +goose Up
ALTER TABLE auth.users ADD COLUMN password_token varchar;
CREATE UNIQUE INDEX idx_users_password_token_unique ON auth.users (password_token);

-- +goose Down
DROP INDEX auth.idx_users_password_token_unique;
ALTER TABLE auth.users DROP COLUMN password_token;
