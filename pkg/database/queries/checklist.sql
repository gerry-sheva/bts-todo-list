-- name: GetChecklist :many
select * from checklist
where user_id = $1 and deleted_at is null;

-- name: CreateChecklist :one
insert into checklist(user_id, title)
values ($1, $2)
returning title;

-- name: GetChecklistDetails :one
select * from checklist
where checklist_id = $1 and user_id = $2 and deleted_at is null;

-- name: DeleteChecklist :exec
update checklist
set deleted_at = now()
where checklist_id = $1 and user_id = $2;
