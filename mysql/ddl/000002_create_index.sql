-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create index account_refresh_token_idx_token on account_refresh_tokens (token);
create index user_idx_email on users (email);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back