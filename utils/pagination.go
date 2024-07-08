package utils

import (
	"strconv"
)

func Paginate(pageStr, limitStr string) (int, int) {
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	return page, limit
}
