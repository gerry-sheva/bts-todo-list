-- +goose Up
-- +goose StatementBegin
create table checklist (
    checklist_id uuid primary key default uuid_generate_v1mc (),
    user_id uuid,
    title text not null,
    created_at timestamptz not null default now (),
    updated_at timestamptz,
    deleted_at timestamptz,
    foreign key (user_id) references users (user_id)
);

select
    trigger_updated_at ('checklist');

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop table if exists checklist;

-- +goose StatementEnd
