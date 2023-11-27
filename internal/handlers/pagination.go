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

func (t PaginationToken) String() string {
	ret, err := MarshalPaginationToken(t)
	if err != nil {
		log.Error().Err(err).Msg("Error marshalling pagination token.")
		return ""
	}
	return ret
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
	log := log.Ctx(r.Context())
	ret := PaginationToken{
		Page:     1,
		PageSize: defaultPageSize,
	}
	var err error

	token := r.URL.Query().Get("cursor")
	if token != "" {
		ret, err = UnmarshalPaginationToken(token)
		if err != nil {
			return PaginationToken{}, err
		}
		log.Debug().
			Str("tokenString", token).
			Int("pageSize", ret.PageSize).
			Int("page", ret.Page).
			Msg("Pagination token found.")
	}

	pageSizeStr := r.URL.Query().Get("pageSize")
	if pageSizeStr != "" {
		ret.PageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil {
			return PaginationToken{}, err
		}
	}

	log.Debug().
		Int("pageSize", ret.PageSize).
		Int("page", ret.Page).
		Msg("Pagination Token Processing Complete.")

	return ret, nil
}

func Paginate[D PageableDB[D]](db D, token *PaginationToken) D {
	if token == nil {
		return db.WithPagination(1, defaultPageSize)
	}
	return db.WithPagination(token.Page, token.PageSize)
}
