-- +goose Up
-- +goose StatementBegin
ALTER TABLE logs RENAME TO logsUser;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE logsUser RENAME TO logs;
-- +goose StatementEnd
