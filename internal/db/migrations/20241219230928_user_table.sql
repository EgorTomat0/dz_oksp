-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.users
(
    id       SERIAL,
    name     VARCHAR(255),
    reg_date DATE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.users;
-- +goose StatementEnd
