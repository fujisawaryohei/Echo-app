-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts (
    id SERIAL NOT NULL,
    title VARCHAR(50) NOT NULL,
    body TEXT NOT NULL,
    published BOOLEAN,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS posts;
-- +goose StatementEnd

