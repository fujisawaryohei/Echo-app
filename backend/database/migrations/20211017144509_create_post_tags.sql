-- +goose Up
-- +goose StatementBegin
CREATE TABLE post_tags (
    id SERIAL NOT NULL,
    post_id INTEGER NOT NULL,
    tag_id INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY(post_id) REFERENCES posts(id),
    FOREIGN KEY(tag_id) REFERENCES tags(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS post_tags;
-- +goose StatementEnd

