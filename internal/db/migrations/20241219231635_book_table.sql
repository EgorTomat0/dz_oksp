-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.book
(
    id       SERIAL,
    title     VARCHAR(255),
    description TEXT,
    release_date DATE,
    rating float,
    cover VARCHAR,
    paragraph TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.book;
-- +goose StatementEnd
