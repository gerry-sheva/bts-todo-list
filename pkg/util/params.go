package util

import (
	"errors"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrInvalidUUID = errors.New("Invalid UUID format")
)

func ParseUUIDParam(r *http.Request, prefix string) (pgtype.UUID, error) {
	path := strings.TrimPrefix(r.URL.Path, prefix)

	uuidStr := strings.Split(path, "/")[0]

	var uuid pgtype.UUID
	err := uuid.Scan(uuidStr)
	if err != nil {
		return pgtype.UUID{}, ErrInvalidUUID
	}

	return uuid, nil
}
