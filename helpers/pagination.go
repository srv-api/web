package helpers

import (
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/srv-api/web/dto"
)

// helpers/pagination.go
func GeneratePaginationRequest(c echo.Context) *dto.Pagination {
	limit := 10
	page := 1
	sort := "merchant_details.created_at desc"

	var searchs []dto.Search
	query := c.QueryParams()

	for key, values := range query {
		val := values[len(values)-1]

		switch key {
		case "limit":
			if v, err := strconv.Atoi(val); err == nil && v > 0 {
				limit = v
			}
		case "page":
			if v, err := strconv.Atoi(val); err == nil && v > 0 {
				page = v
			}
		case "sort":
			// ⛔ whitelist biar ga SQL injection
			if strings.Contains(val, "created_at") {
				sort = val
			}
		default:
			if strings.Contains(key, ".") {
				s := strings.Split(key, ".")
				searchs = append(searchs, dto.Search{
					Column: s[0],
					Action: s[1],
					Query:  val,
				})
			}
		}
	}

	return &dto.Pagination{
		Limit:   limit,
		Page:    page,
		Sort:    sort,
		Searchs: searchs,
	}
}

func TruncateString(input string, length int) string {
	if len(input) > length {
		return input[:length] + "..."
	}
	return input
}
