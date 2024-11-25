-- +goose Up
-- +goose StatementBegin
ALTER TABLE client ADD COLUMN hash_password VARCHAR(255) SET NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE client DROP COLUMN hash_password;
-- +goose StatementEnd
