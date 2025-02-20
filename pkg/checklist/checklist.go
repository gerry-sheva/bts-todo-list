package checklist

import (
	"context"

	"github.com/gerry-sheva/bts-todo-list/pkg/database/repository"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func createChecklist(ctx context.Context, dbpool *pgxpool.Pool, user_id pgtype.UUID, i *CreateChecklistInput) (string, error) {
	p := repository.CreateChecklistParams{
		UserID: user_id,
		Title:  i.Title,
	}
	checklist, err := repository.New(dbpool).CreateChecklist(ctx, p)
	if err != nil {
		return "", err
	}

	return checklist, nil
}

func deleteChecklist(ctx context.Context, dbpool *pgxpool.Pool, user_id pgtype.UUID, checklist_id pgtype.UUID) error {
	p := repository.DeleteChecklistParams{
		UserID:      user_id,
		ChecklistID: checklist_id,
	}

	err := repository.New(dbpool).DeleteChecklist(ctx, p)
	if err != nil {
		return err
	}

	return nil
}

func getAllChecklist(ctx context.Context, dbpool *pgxpool.Pool, user_id pgtype.UUID) ([]repository.Checklist, error) {
	checklists, err := repository.New(dbpool).GetChecklist(ctx, user_id)
	if err != nil {
		return nil, err
	}

	return checklists, nil
}
