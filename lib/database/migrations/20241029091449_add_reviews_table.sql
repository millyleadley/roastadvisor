-- +goose Up
-- +goose StatementBegin
CREATE TABLE reviews (
    id text PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    user_id text REFERENCES users(id) NOT NULL,
    comment text NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    archived_at TIMESTAMPTZ NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE reviews;
-- +goose StatementEnd
