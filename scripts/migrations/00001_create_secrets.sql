-- +goose Up
-- +goose StatementBegin
create table if not exists secrets
(
    id       bigserial primary key,
    login    text not null,
    password text not null,
    role     text not null default 'employee'
);
-- +goose StatementEnd
-- +goose StatementBegin
create unique index unique_login on secrets (login);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists secrets;
-- +goose StatementEnd
-- +goose StatementBegin
drop index if exists unique_login;
-- +goose StatementEnd