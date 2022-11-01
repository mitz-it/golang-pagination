package pagination

import (
	"github.com/gin-gonic/gin"
	paginate "github.com/gobeam/mongo-go-pagination"
)

type IPaginationProvider interface {
	GetPaginationQuery(ctx *gin.Context) *PaginationQuery
	GetPaginationResult(ctx *gin.Context) *PaginationResult
	SetPaginationData(paginationData *paginate.PaginationData)
}
