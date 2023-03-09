-- +goose Up
-- +goose StatementBegin
create table if not exists secrets
(
    id       bigserial primary key,
    login    text not null,
    password text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists secrets;
-- +goose StatementEnd
