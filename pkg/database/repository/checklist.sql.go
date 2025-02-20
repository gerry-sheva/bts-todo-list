// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: checklist.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createChecklist = `-- name: CreateChecklist :one
insert into checklist(user_id, title)
values ($1, $2)
returning title
`

type CreateChecklistParams struct {
	UserID pgtype.UUID
	Title  string
}

func (q *Queries) CreateChecklist(ctx context.Context, arg CreateChecklistParams) (string, error) {
	row := q.db.QueryRow(ctx, createChecklist, arg.UserID, arg.Title)
	var title string
	err := row.Scan(&title)
	return title, err
}

const deleteChecklist = `-- name: DeleteChecklist :exec
update checklist
set deleted_at = now()
where checklist_id = $1 and user_id = $2
`

type DeleteChecklistParams struct {
	ChecklistID pgtype.UUID
	UserID      pgtype.UUID
}

func (q *Queries) DeleteChecklist(ctx context.Context, arg DeleteChecklistParams) error {
	_, err := q.db.Exec(ctx, deleteChecklist, arg.ChecklistID, arg.UserID)
	return err
}

const getChecklist = `-- name: GetChecklist :many
select checklist_id, user_id, title, created_at, updated_at, deleted_at from checklist
where user_id = $1 and deleted_at is null
`

func (q *Queries) GetChecklist(ctx context.Context, userID pgtype.UUID) ([]Checklist, error) {
	rows, err := q.db.Query(ctx, getChecklist, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Checklist
	for rows.Next() {
		var i Checklist
		if err := rows.Scan(
			&i.ChecklistID,
			&i.UserID,
			&i.Title,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getChecklistDetails = `-- name: GetChecklistDetails :one
select
    c.title,
    ci.item,
    ci.checked_at
from checklist c
join checklist_items ci on c.checklist_id = ci.checklist_id
where c.checklist_id = $1 and c.user_id = $2 and c.deleted_at is null
`

type GetChecklistDetailsParams struct {
	ChecklistID pgtype.UUID
	UserID      pgtype.UUID
}

type GetChecklistDetailsRow struct {
	Title     string
	Item      string
	CheckedAt pgtype.Timestamptz
}

func (q *Queries) GetChecklistDetails(ctx context.Context, arg GetChecklistDetailsParams) (GetChecklistDetailsRow, error) {
	row := q.db.QueryRow(ctx, getChecklistDetails, arg.ChecklistID, arg.UserID)
	var i GetChecklistDetailsRow
	err := row.Scan(&i.Title, &i.Item, &i.CheckedAt)
	return i, err
}
