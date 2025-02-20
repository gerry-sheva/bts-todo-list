-- name: CreateChecklistItem :one
insert into checklist_items(checklist_id, item)
values($1, $2) returning item;

-- name: CheckChecklistItem :exec
update checklist_items
set checked_at = now()
where checklist_item_id = $1;

-- name: DeleteChecklistItem :exec
update checklist_items
set deleted_at = now()
where checklist_item_id = $1;

-- name: UpdateChecklistItem :one
update checklist_items
set item = $1
where checklist_item_id = $2
returning item;
