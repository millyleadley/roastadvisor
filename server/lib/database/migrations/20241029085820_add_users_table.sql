-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id text PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    email text NOT NULL UNIQUE, 
    hashed_password text NOT NULL,
    session_token text,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    archived_at TIMESTAMPTZ NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
