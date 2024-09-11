package Util

import (
	"math"
)

// PaginatedList represents a paginated list of items
type PaginatedList[T any] struct {
	Items      []T
	TotalCount int64
	PageIndex  int
	PageSize   int
	TotalPages int
}

// NewPaginatedList creates a new PaginatedList
func NewPaginatedList[T any](items []T, count int64, pageIndex, pageSize int) PaginatedList[T] {
	totalPages1 := int(math.Ceil(float64(count) / float64(pageSize)))
	return PaginatedList[T]{
		Items:      items,
		TotalCount: count,
		PageIndex:  pageIndex,
		PageSize:   pageSize,
		TotalPages: totalPages1,
	}
}
