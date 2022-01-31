package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"idist-go/app/providers/configProvider"
	"strings"
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
	config := configProvider.GetConfig()
	if request.Page <= 0 {
		request.Page = 1
	}
	if request.Length <= 0 {
		request.Length = 12
	}
	if request.Search != "" {
		request.Search = strings.TrimSpace(request.Search)
	}
	if request.Length > config.GetInt64("pagination.max_length") {
		request.Length = config.GetInt64("pagination.max_length")
	}
	if request.Length < config.GetInt64("pagination.min_length") {
		request.Length = config.GetInt64("pagination.min_length")
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

func ResponseSuccess(c *gin.Context, code int, message string, data interface{}) {
	con := configProvider.GetConfig()
	c.AbortWithStatusJSON(code, gin.H{
		"code":    code,
		"data":    data,
		"message": message,
		"version": con.GetString("app.server.version"),
	})
	panic(nil)
}

func ResponseError(c *gin.Context, code int, message string, err error) {
	con := configProvider.GetConfig()
	if con.GetString("app.env") == "production" {
		err = nil
	}
	if err == nil {
		err = errors.New("")
	}
	c.AbortWithStatusJSON(code, gin.H{
		"code":    code,
		"error":   err.Error(),
		"message": message,
		"version": con.GetString("app.server.version"),
	})
	panic(nil)
}

type Pagination struct {
	Page      int64  `json:"page" form:"page"`
	Length    int64  `json:"length" form:"length"`
	Search    string `json:"search" form:"search"`
	Order     string `json:"order_by" form:"order_by"`
	Dir       string `json:"order_dir" form:"order_dir"`
	Total     int64  `json:"total"`
	DirNumber int    `json:"-"`
	Deleted   bool   `json:"-"`
}

func BindPagination(c *gin.Context) Pagination {
	var request Pagination
	_ = c.BindQuery(&request)
	config := configProvider.GetConfig()
	if request.Page <= 0 {
		request.Page = 1
	}
	if request.Length <= 0 {
		request.Length = 12
	}
	if request.Search != "" {
		request.Search = strings.TrimSpace(request.Search)
	}
	if request.Length > config.GetInt64("pagination.max_length") {
		request.Length = config.GetInt64("pagination.max_length")
	}
	if request.Length < config.GetInt64("pagination.min_length") {
		request.Length = config.GetInt64("pagination.min_length")
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
