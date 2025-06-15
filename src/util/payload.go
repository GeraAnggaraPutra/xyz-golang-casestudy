package util

import (
	"fmt"

	"kredit-plus/src/constant"
)

type PaginationPayload struct {
	Search      string `query:"search" json:"search"`
	Sort        string `query:"sort" json:"sort"`
	Direction   string `query:"direction" json:"direction"`
	Page        int    `query:"page" json:"page"`
	Limit       int    `query:"limit" json:"limit"`
	SetSearch   bool
	Offset      int
	Order       string
	SetPaginate bool
}

// Initialize pagination payload.
func (p *PaginationPayload) Init() {
	if p.Search != "" {
		p.SetSearch = true
		p.Search = fmt.Sprintf("%%%s%%", p.Search)
	}

	if p.Sort == "" || p.Direction == "" {
		p.Order = constant.DefaultOrder
	} else {
		p.Order = fmt.Sprintf("%s %s", p.Sort, p.Direction)
	}

	if !(p.Page <= 0 && p.Limit <= 0) {
		if p.Page <= 0 {
			p.Page = constant.DefaultPage
		}

		if p.Limit <= 0 {
			p.Limit = constant.DefaultLimit
		}

		p.Offset = (p.Page * p.Limit) - p.Limit
		p.SetPaginate = true
	}
}
