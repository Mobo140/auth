-- +goose Up
-- +goose StatementBegin
ALTER TABLE access ADD COLUMN updated_at TIMESTAMP;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE access DROP COLUMN updated_at;
-- +goose StatementEnd
