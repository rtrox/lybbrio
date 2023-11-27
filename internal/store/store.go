package store

type Store interface {
	WithPagination(page, pageSize int) Store
}
