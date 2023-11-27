package handlers

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

const defaultPageSize = 30

type paginationCtxKeyType string

const paginationCtxKey paginationCtxKeyType = "pagination"

type PaginationContextObject struct {
	Token    PaginationToken
	Response PaginationResponse
}

func (o PaginationContextObject) NextCursor() string {
	return o.Token.Next().String()
}

type PaginationToken struct {
	// TODO: implement true cursors
	Page     int
	PageSize int
}

type PaginationResponse struct {
	NextURL    string `json:"next"`
	NextCursor string `json:"nextCursor"`
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

func PaginationCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := PaginationTokenFromRequest(r)
		ctx := r.Context()
		log := log.Ctx(ctx)
		if err != nil {
			log.Error().Err(err).Msg("Error processing pagination token.")
			render.Render(w, r, ErrBadRequest(err))
			return
		}
		nextToken, err := MarshalPaginationToken(token.Next())
		if err != nil {
			render.Render(w, r, ErrInternalError(AppError{ErrPaginationToken, err.Error()}))
			return
		}
		obj := PaginationContextObject{
			Token: token,
			Response: PaginationResponse{
				NextCursor: nextToken,
				NextURL:    fmt.Sprintf("%s?cursor=%s", r.URL.Path, nextToken),
			},
		}
		ctx = context.WithValue(ctx, paginationCtxKey, obj)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func PaginationCtxFromRequest(r *http.Request) PaginationContextObject {
	return r.Context().Value(paginationCtxKey).(PaginationContextObject)
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

func Paginate[D PageableDB[D]](db D, token PaginationToken) D {
	return db.WithPagination(token.Page, token.PageSize)
}
