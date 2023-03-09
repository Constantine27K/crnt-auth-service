-- +goose Up
-- +goose StatementBegin
create table if not exists contacts
(
    id           bigserial primary key,
    phone_number text default '',
    email        text default '',
    telegram_url text default '',
    discord_url  text default ''
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists contacts;
-- +goose StatementEnd
