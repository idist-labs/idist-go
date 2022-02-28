package helpers

import (
	"regexp"
	"strings"
)

func MakeSlug(name string) string {
	/* Chuyển sang chuỗi chữ thường */
	name = strings.ToLower(name)

	/* Liệt kê danh sách các mẫu regex */
	listRegex := []struct {
		Regex   string
		Replace string
	}{
		{Regex: "à|á|ạ|ả|ã|â|ầ|ấ|ậ|ẩ|ẫ|ă|ằ|ắ|ặ|ẳ|ẵ", Replace: "a"},
		{Regex: "è|é|ẹ|ẻ|ẽ|ê|ề|ế|ệ|ể|ễ", Replace: "e"},
		{Regex: "ì|í|ị|ỉ|ĩ", Replace: "i"},
		{Regex: "ò|ó|ọ|ỏ|õ|ô|ồ|ố|ộ|ổ|ỗ|ơ|ờ|ớ|ợ|ở|ỡ", Replace: "o"},
		{Regex: "ù|ú|ụ|ủ|ũ|ư|ừ|ứ|ự|ử|ữ", Replace: "u"},
		{Regex: "ỳ|ý|ỵ|ỷ|ỹ", Replace: "y"},
		{Regex: "đ", Replace: "d"},
		{Regex: "\u0300|\u0301|\u0303|\u0309|\u0323", Replace: ""},
		{Regex: "\u02C6|\u0306|\u031B", Replace: ""},
		{Regex: "\u02C6|\u0306|\u031B", Replace: ""},
		{Regex: `\s+`, Replace: " "},
		{Regex: `[^a-z0-9]+`, Replace: "-"},
	}

	/* Duyệt danh sách và thay thế */
	for _, item := range listRegex {
		if r, err := regexp.Compile(item.Regex); err == nil {
			name = r.ReplaceAllString(name, item.Replace)
		}
	}
	return strings.Trim(name, " ")
}

func MakeAlias(name string) string {
	/* Chuyển sang chuỗi chữ thường */
	name = strings.ToLower(name)

	/* Liệt kê danh sách các mẫu regex */
	listRegex := []struct {
		Regex   string
		Replace string
	}{
		{Regex: "à|á|ạ|ả|ã|â|ầ|ấ|ậ|ẩ|ẫ|ă|ằ|ắ|ặ|ẳ|ẵ", Replace: "a"},
		{Regex: "è|é|ẹ|ẻ|ẽ|ê|ề|ế|ệ|ể|ễ", Replace: "e"},
		{Regex: "ì|í|ị|ỉ|ĩ", Replace: "i"},
		{Regex: "ò|ó|ọ|ỏ|õ|ô|ồ|ố|ộ|ổ|ỗ|ơ|ờ|ớ|ợ|ở|ỡ", Replace: "o"},
		{Regex: "ù|ú|ụ|ủ|ũ|ư|ừ|ứ|ự|ử|ữ", Replace: "u"},
		{Regex: "ỳ|ý|ỵ|ỷ|ỹ", Replace: "y"},
		{Regex: "đ", Replace: "d"},
		{Regex: "\u0300|\u0301|\u0303|\u0309|\u0323", Replace: ""},
		{Regex: "\u02C6|\u0306|\u031B", Replace: ""},
		{Regex: "\u02C6|\u0306|\u031B", Replace: ""},
		{Regex: `\s+`, Replace: " "},
	}

	/* Duyệt danh sách và thay thế */
	for _, item := range listRegex {
		if r, err := regexp.Compile(item.Regex); err == nil {
			name = r.ReplaceAllString(name, item.Replace)
		}
	}
	return strings.Trim(name, " ")
}
