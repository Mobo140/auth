-- +goose Up
-- +goose StatementBegin
CREATE TABLE access(
    id SERIAL PRIMARY KEY,
    endpoint VARCHAR(255) NOT NULL,
    role INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE access;
-- +goose StatementEnd
