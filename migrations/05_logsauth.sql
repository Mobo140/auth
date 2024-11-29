-- +goose Up
-- +goose StatementBegin
CREATE TABLE logsAuth(
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    activity VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE logs;
-- +goose StatementEnd
