-- +goose Up
-- +goose StatementBegin
create table users (
    user_id uuid primary key default uuid_generate_v1mc (),
    username text collate "case_insensitive" unique not null,
    password text not null,
    created_at timestamptz not null default now ()
);

-- +goose StatementEnd
--
-- +goose Down
-- +goose StatementBegin
drop table if exists users;

-- +goose StatementEnd
