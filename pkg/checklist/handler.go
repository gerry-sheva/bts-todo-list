package checklist

import (
	"log/slog"
	"net/http"

	"github.com/gerry-sheva/bts-todo-list/pkg/apierror"
	"github.com/gerry-sheva/bts-todo-list/pkg/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ChecklistHandler struct {
	logger *slog.Logger
	dbpool *pgxpool.Pool
}

func New(logger *slog.Logger, dbpool *pgxpool.Pool) *ChecklistHandler {
	return &ChecklistHandler{
		logger,
		dbpool,
	}
}

func (h *ChecklistHandler) CreateChecklist(w http.ResponseWriter, r *http.Request) {
	var i CreateChecklistInput
	if err := util.ReadJSON(w, r, &i); err != nil {
		apierror.GlobalErrorHandler.BadRequestResponse(w, r, err)
		return
	}

	h.logger.Info("Creating new checklist")

	v := util.NewValidator()

	if i.validate(v); !v.Valid() {
		apierror.GlobalErrorHandler.FailedValidationResponse(w, r, v.Errors)
		return
	}

	var user_id pgtype.UUID
	user_id.Scan(r.Context().Value("sub").(string))

	checklist, err := createChecklist(r.Context(), h.dbpool, user_id, &i)
	if err != nil {
		apierror.GlobalErrorHandler.ServerErrorResponse(w, r, err)
		return
	}

	err = util.WriteJSON(w, http.StatusOK, util.Envelope{"checklist": checklist}, nil)
	if err != nil {
		apierror.GlobalErrorHandler.ServerErrorResponse(w, r, err)
	}
}

func (h *ChecklistHandler) DeleteChecklist(w http.ResponseWriter, r *http.Request) {
	var user_id pgtype.UUID
	user_id.Scan(r.Context().Value("sub").(string))

	var checklist_id pgtype.UUID
	checklist_id.Scan(r.PathValue("checklist_id"))

	err := deleteChecklist(r.Context(), h.dbpool, user_id, checklist_id)
	if err != nil {
		apierror.GlobalErrorHandler.ServerErrorResponse(w, r, err)
	}
}

func (h *ChecklistHandler) GetAllChecklist(w http.ResponseWriter, r *http.Request) {
	var user_id pgtype.UUID
	user_id.Scan(r.Context().Value("sub").(string))

	checklists, err := getAllChecklist(r.Context(), h.dbpool, user_id)
	if err != nil {
		apierror.GlobalErrorHandler.NotFoundResponse(w, r)
	}

	err = util.WriteJSON(w, http.StatusOK, util.Envelope{"checklists": checklists}, nil)
	if err != nil {
		apierror.GlobalErrorHandler.ServerErrorResponse(w, r, err)
	}
}

func (h *ChecklistHandler) GetChecklistDetails(w http.ResponseWriter, r *http.Request) {
	var user_id pgtype.UUID
	user_id.Scan(r.Context().Value("sub").(string))

	var checklist_id pgtype.UUID
	checklist_id.Scan(r.PathValue("checklist_id"))

	checklist, err := getChecklistDetails(r.Context(), h.dbpool, user_id, checklist_id)
	if err != nil {
		apierror.GlobalErrorHandler.ServerErrorResponse(w, r, err)
	}

	err = util.WriteJSON(w, http.StatusOK, util.Envelope{"checklist": checklist}, nil)
	if err != nil {
		apierror.GlobalErrorHandler.ServerErrorResponse(w, r, err)
	}
}
