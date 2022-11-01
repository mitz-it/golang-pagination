package pagination

import "errors"

type SortOrientation int

const (
	ASC  SortOrientation = 1
	DESC SortOrientation = -1
)

func (sort SortOrientation) Int() int {
	return int(sort)
}

func ToSortOrientation(sort string) SortOrientation {
	if sort == "ASC" {
		return ASC
	} else if sort == "DESC" {
		return DESC
	}
	err := errors.New("could not parse sort")
	panic(err)
}

func (sort SortOrientation) String() string {
	if sort == ASC {
		return "ASC"
	} else {
		return "DESC"
	}
}

type PaginationQuery struct {
	Size   int64           `json:"size"`
	Page   int64           `json:"page"`
	Sort   SortOrientation `json:"sort"`
	SortBy string          `json:"sortBy"`
}

type PaginationResult struct {
	BaseURL    string `json:"baseURL"`
	PageURL    string `json:"pageURL"`
	PrevURL    string `json:"prevURL"`
	NextURL    string `json:"nextURL"`
	Total      int64  `json:"total"`
	Page       int64  `json:"page"`
	PerPage    int64  `json:"perPage"`
	Prev       int64  `json:"prev"`
	Next       int64  `json:"next"`
	TotalPages int64  `json:"totalPages"`
}
