package pagination

import (
	"fmt"

	"github.com/gin-gonic/gin"
	paginate "github.com/gobeam/mongo-go-pagination"
)

type MongoDbPaginationProvider struct {
	paginationData  *paginate.PaginationData
	paginationQuery *PaginationQuery
}

func (provider *MongoDbPaginationProvider) GetPaginationQuery(ctx *gin.Context) *PaginationQuery {
	paginationQuery := buildPaginationQuery(ctx)
	provider.paginationQuery = paginationQuery
	return paginationQuery
}

func (provider *MongoDbPaginationProvider) GetPaginationResult(ctx *gin.Context) *PaginationResult {
	baseURL := provider.getBaseURL(ctx)
	pageURL := provider.getPageURL(baseURL)
	prevURL := provider.getPrevURL(baseURL)
	nextURL := provider.getNextURL(baseURL)
	return &PaginationResult{
		BaseURL:    baseURL,
		PageURL:    pageURL,
		PrevURL:    prevURL,
		NextURL:    nextURL,
		Total:      provider.paginationData.Total,
		Page:       provider.paginationData.Page,
		PerPage:    provider.paginationData.PerPage,
		Prev:       provider.paginationData.Prev,
		Next:       provider.paginationData.Next,
		TotalPages: provider.paginationData.TotalPage,
	}
}

func (provider *MongoDbPaginationProvider) SetPaginationData(paginationData *paginate.PaginationData) {
	provider.paginationData = paginationData
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

func (provider *MongoDbPaginationProvider) getPrevURL(baseURL string) string {
	prevURL := provider.getBaseQueryString(baseURL, provider.paginationData.Prev)

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
