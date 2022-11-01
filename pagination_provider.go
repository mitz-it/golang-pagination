package pagination

import (
	"github.com/gin-gonic/gin"
)

type IPaginationProvider interface {
	GetPaginationQuery(ctx *gin.Context) *PaginationQuery
	GetPaginationResult(ctx *gin.Context) *PaginationResult
	SetPaginationData(paginationData *PaginationData)
}
