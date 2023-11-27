package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

const defaultPageSize = 30

type PaginationToken struct {
	// TODO: implement true cursors
	Page     int
	PageSize int
}

type PageableDB[D any] interface {
	WithPagination(page, pageSize int) D
}

func UnmarshalPaginationToken(s string) (PaginationToken, error) {
	// buf := strings.NewReader(s)
	// decoder := base64.NewDecoder(base64.URLEncoding, buf)
	// ret := PaginationToken{}
	// err := gob.NewDecoder(decoder).Decode(&ret)
	ret := PaginationToken{}
	err := json.NewDecoder(base64.NewDecoder(base64.URLEncoding, strings.NewReader(s))).Decode(&ret)
	return ret, err
}

func MarshalPaginationToken(t PaginationToken) (string, error) {
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.URLEncoding, &buf)
	defer encoder.Close()

	// err := gob.NewEncoder(encoder).Encode(t)
	err := json.NewEncoder(encoder).Encode(t)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (t PaginationToken) Next() PaginationToken {
	return PaginationToken{
		Page:     t.Page + 1,
		PageSize: t.PageSize,
	}
}

func PaginationTokenFromRequest(r *http.Request) (PaginationToken, error) {
	token := r.URL.Query().Get("cursor")
	log.Debug().Str("cursor", token).Msg("cursor")
	if token != "" {
		return UnmarshalPaginationToken(token)
	}

	pageSizeStr := r.URL.Query().Get("pageSize")
	pageSize := defaultPageSize
	if pageSizeStr != "" {
		var err error
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil {
			return PaginationToken{}, err
		}
	}
	return PaginationToken{1, pageSize}, nil
}

func Paginate[D PageableDB[D]](db D, token *PaginationToken) D {
	if token == nil {
		return db.WithPagination(1, defaultPageSize)
	}
	return db.WithPagination(token.Page, token.PageSize)
}
