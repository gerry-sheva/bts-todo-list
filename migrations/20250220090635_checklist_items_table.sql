-- +goose Up
-- +goose StatementBegin
create table checklist_items (
    checklist_item_id uuid primary key default uuid_generate_v1mc (),
    checklist_id uuid,
    item text not null,
    checked_at timestamptz,
    created_at timestamptz not null default now (),
    updated_at timestamptz,
    deleted_at timestamptz,
    foreign key (checklist_id) references checklist (checklist_id)
);

select
    trigger_updated_at ('checklist_items');

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop table if exists checklist_items;

-- +goose StatementEnd
