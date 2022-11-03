package pagination

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getPage(ctx *gin.Context) (int64, error) {
	page_str := ctx.Request.URL.Query().Get("page")
	page, err := strconv.Atoi(page_str)

	if err != nil {
		return defaultPage, err
	}

	return int64(page), nil
}

func getSize(ctx *gin.Context) (int64, error) {
	size_str := ctx.Request.URL.Query().Get("size")
	size, err := strconv.Atoi(size_str)

	if err != nil {
		return defaultSize, err
	}

	return int64(size), nil
}

func getSort(ctx *gin.Context) (*SortOrientation, error) {
	sort := ctx.Request.URL.Query().Get("sort")

	if sort == "" {
		err := errors.New("sort parameter was not in the query string")
		return nil, err
	}

	sortOrientation := ToSortOrientation(sort)

	return &sortOrientation, nil
}

func getSortBy(ctx *gin.Context) string {
	sortBy := ctx.Request.URL.Query().Get("sortBy")
	return sortBy
}

func buildPaginationQuery(ctx *gin.Context, fallbacks []FallBackPaginationFunc) *PaginationQuery {
	paginationQuery := new(PaginationQuery)

	for _, fallback := range fallbacks {
		fallback(paginationQuery)
	}

	page, err := getPage(ctx)

	if err == nil {
		paginationQuery.Page = page
	}

	size, err := getSize(ctx)

	if err == nil {
		paginationQuery.Size = size
	}

	sort, err := getSort(ctx)

	if *sort == ASC || *sort == DESC {
		paginationQuery.Sort = *sort
	} else if err != nil && sort == nil && paginationQuery.Sort == 0 {
		paginationQuery.Sort = defaultSort
	}

	sortBy := getSortBy(ctx)

	paginationQuery.SortBy = sortBy

	return paginationQuery
}
