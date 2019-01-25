package paging

import (
	"math"
)

/*
PagingInfo describes common information used in paging
*/
type PagingInfo struct {
	CurrentPage     int  `json:"currentPage"`
	End             int  `json:"end"`
	HasNextPage     bool `json:"hasNextPage"`
	HasPreviousPage bool `json:"hasPreviousPage"`
	NextPage        int  `json:"nextPage"`
	PageSize        int  `json:"pageSize"`
	PreviousPage    int  `json:"previousPage"`
	Start           int  `json:"start"`
	TotalItems      int  `json:"totalItems"`
	TotalPages      int  `json:"totalPages"`
}

/*
Calculate determines our paging information. This is done by using
the current page we are on and page size vs. the total items
available. Using this information we can determine is we have more
pages
*/
func (p *PagingInfo) Calculate(currentPage, pageSize, totalItems int) {
	p.TotalItems = totalItems
	p.PageSize = pageSize
	p.CurrentPage = currentPage

	if p.PageSize <= 0 {
		p.PageSize = totalItems
	}

	p.TotalPages = int(math.Ceil(float64(totalItems / p.PageSize)))
	p.NextPage = currentPage + 1
	p.PreviousPage = currentPage - 1
	p.HasNextPage = true
	p.HasPreviousPage = true

	if p.TotalPages < 1 {
		p.TotalPages = 1
	}

	if p.NextPage > p.TotalPages {
		p.NextPage = p.TotalPages
		p.HasNextPage = false
	}

	if p.PreviousPage < 1 {
		p.PreviousPage = 1
		p.HasPreviousPage = false
	}

	p.Start = (p.CurrentPage - 1) * p.PageSize
	if p.Start > totalItems {
		p.Start = 0
	}

	p.End = p.Start + p.PageSize
	if p.End > totalItems {
		p.End = totalItems
	}
}
