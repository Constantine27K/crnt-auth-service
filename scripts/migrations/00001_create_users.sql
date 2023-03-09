-- +goose Up
-- +goose StatementBegin
create table if not exists users
(
    id            bigserial primary key,
    name          text               default '',
    last_name     text               default '',
    display_name  text      not null default '',
    birthday      timestamp,
    employed_at   timestamp,
    fired_at      timestamp,
    about_info    text               default '',
    avatar_url    text      not null default '',
    contacts_id   bigserial,
    salary        double precision   default 0,
    is_piece_wage boolean   not null default false,
    secrets_id    bigserial not null
);
-- +goose StatementEnd
-- +goose StatementBegin
create unique index unique_display_name on users (display_name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
-- +goose StatementBegin
drop index if exists unique_display_name;
-- +goose StatementEnd
