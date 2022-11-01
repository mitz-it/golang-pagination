package pagination

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func getPage(ctx *gin.Context) int64 {
	page_str := ctx.Request.URL.Query().Get("page")
	page, err := strconv.Atoi(page_str)

	panicOnError(err)

	return int64(page)
}

func getSize(ctx *gin.Context) int64 {
	size_str := ctx.Request.URL.Query().Get("size")
	size, err := strconv.Atoi(size_str)

	panicOnError(err)

	return int64(size)
}

func getSort(ctx *gin.Context) SortOrientation {
	sort := ctx.Request.URL.Query().Get("sort")

	if sort == "" {
		sort = "ASC"
	}

	return ToSortOrientation(sort)
}

func getSortBy(ctx *gin.Context) string {
	sortBy := ctx.Request.URL.Query().Get("sortBy")
	return sortBy
}

func buildPaginationQuery(ctx *gin.Context) *PaginationQuery {
	page := getPage(ctx)
	size := getSize(ctx)
	sort := getSort(ctx)
	sortBy := getSortBy(ctx)
	return &PaginationQuery{
		Page:   page,
		Size:   size,
		Sort:   sort,
		SortBy: sortBy,
	}
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
