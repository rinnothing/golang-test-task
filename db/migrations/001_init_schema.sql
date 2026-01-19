-- +goose Up

CREATE TABLE integers (
    integer BIGINT
);

-- +goose Down
DROP TABLE integers;
