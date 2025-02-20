-- name: GetChecklist :many
select * from checklist
where user_id = $1 and deleted_at is null;

-- name: CreateChecklist :one
insert into checklist(user_id, title)
values ($1, $2)
returning title;

-- name: GetChecklistDetails :one
select
    c.title,
    ci.item,
    ci.checked_at
from checklist c
join checklist_items ci on c.checklist_id = ci.checklist_id
where c.checklist_id = $1 and c.user_id = $2 and c.deleted_at is null;

-- name: DeleteChecklist :exec
update checklist
set deleted_at = now()
where checklist_id = $1 and user_id = $2;
