package pagination

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type MongoDbPaginationProvider struct {
	paginationData  *PaginationData
	paginationQuery *PaginationQuery
}

func (provider *MongoDbPaginationProvider) GetPaginationQuery(ctx *gin.Context) *PaginationQuery {
	paginationQuery := buildPaginationQuery(ctx)
	provider.paginationQuery = paginationQuery
	return paginationQuery
}

func (provider *MongoDbPaginationProvider) GetPaginationResult(ctx *gin.Context) *PaginationResult {
	prev := provider.getPrev()
	baseURL := provider.getBaseURL(ctx)
	pageURL := provider.getPageURL(baseURL)
	prevURL := provider.getPrevURL(baseURL, prev)
	nextURL := provider.getNextURL(baseURL)
	return &PaginationResult{
		BaseURL:    baseURL,
		PageURL:    pageURL,
		PrevURL:    prevURL,
		NextURL:    nextURL,
		Total:      provider.paginationData.Total,
		Page:       provider.paginationData.Page,
		PerPage:    provider.paginationData.PerPage,
		Prev:       prev,
		Next:       provider.paginationData.Next,
		TotalPages: provider.paginationData.TotalPage,
	}
}

func (provider *MongoDbPaginationProvider) SetPaginationData(paginationData *PaginationData) {
	provider.paginationData = paginationData
}

func (provider *MongoDbPaginationProvider) getPrev() int64 {
	prev := provider.paginationData.Prev

	if prev == 0 {
		prev = 1
	}

	return prev
}

func (provider *MongoDbPaginationProvider) getBaseURL(ctx *gin.Context) string {
	host := ctx.Request.Host
	path := ctx.Request.URL.Path
	baseURL := fmt.Sprintf("%s%s", host, path)
	return baseURL
}

func (provider *MongoDbPaginationProvider) getPageURL(baseURL string) string {
	pageURL := provider.getBaseQueryString(baseURL, provider.paginationData.Page)

	pageURL = provider.appendSortBy(pageURL)

	return pageURL
}

func (provider *MongoDbPaginationProvider) getPrevURL(baseURL string, prev int64) string {
	prevURL := provider.getBaseQueryString(baseURL, prev)

	prevURL = provider.appendSortBy(prevURL)

	return prevURL
}

func (provider *MongoDbPaginationProvider) getNextURL(baseURL string) string {
	nextURL := provider.getBaseQueryString(baseURL, provider.paginationData.Next)

	nextURL = provider.appendSortBy(nextURL)

	return nextURL
}

func (provider *MongoDbPaginationProvider) getBaseQueryString(baseURL string, size int64) string {
	baseQuery := fmt.Sprintf("%s?page=%d&size=%d&sort=%s", baseURL, size, provider.paginationData.PerPage, provider.paginationQuery.Sort.String())
	return baseQuery
}

func (provider *MongoDbPaginationProvider) appendSortBy(url string) string {
	if provider.paginationQuery.SortBy != "" {
		url = fmt.Sprintf("%s&sortBy=%s", url, provider.paginationQuery.SortBy)
	}
	return url
}

func NewMongoDbPaginationProvider() IPaginationProvider {
	return &MongoDbPaginationProvider{
		paginationData:  nil,
		paginationQuery: nil,
	}
}
