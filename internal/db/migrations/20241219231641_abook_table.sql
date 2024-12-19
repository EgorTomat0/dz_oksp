-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.abook
(
    id       SERIAL,
    title     VARCHAR(255),
    description TEXT,
    release_date DATE,
    rating float,
    cover VARCHAR,
    audio VARCHAR
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.abook;
-- +goose StatementEnd