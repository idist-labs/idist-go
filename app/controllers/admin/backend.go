package admin

import (
	"idist-core/app/providers/configProvider"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RequestTable struct {
	Page      int64  `json:"page" form:"page"`
	Length    int64  `json:"length" form:"length"`
	Search    string `json:"search" form:"search"`
	Order     string `json:"order_by" form:"order_by"`
	Dir       string `json:"order_dir" form:"order_dir"`
	DirNumber int    `json:"-"`
	Total     int64  `json:"total"`
	Deleted   bool   `json:"deleted"`
}

func (u *RequestTable) CustomOptions(opts *options.FindOptions) *options.FindOptions {
	return opts.SetSort(bson.M{u.Order: u.DirNumber}).SetLimit(u.Length).SetSkip((u.Page - 1) * u.Length)
}

func (u *RequestTable) CustomFilters(filter bson.M) bson.M {
	if u.Deleted == false {
		filter["deleted_at"] = nil
	}

	return filter
}
func BindRequestTable(c *gin.Context, order string) RequestTable {
	var request RequestTable
	_ = c.BindQuery(&request)
	if request.Page <= 0 {
		request.Page = 1
	}
	if request.Length <= 0 {
		request.Length = 10
	}
	if request.Search != "" {
		request.Search = strings.TrimSpace(request.Search)
	}
	if request.Length > configProvider.GetConfig().GetInt64("pagination.max_length") {
		request.Length = configProvider.GetConfig().GetInt64("pagination.max_length")
	}
	if request.Length < configProvider.GetConfig().GetInt64("pagination.min_length") {
		request.Length = configProvider.GetConfig().GetInt64("pagination.min_length")
	}
	if request.Order == "" {
		request.Order = order
	}
	if strings.ToLower(request.Dir) != "asc" {
		request.DirNumber = -1
		request.Dir = "desc"
	} else {
		request.DirNumber = 1
		request.Dir = "asc"
	}
	return request
}
