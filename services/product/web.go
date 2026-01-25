package product

import (
	"fmt"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/web/dto"
)

// services/product/web.go
func (s *productService) Web(ctx echo.Context, req *dto.Pagination) dto.Response {
	if req.Page < 1 {
		req.Page = 1
	}

	repoResult, totalPages := s.Repo.Web(req)

	if repoResult.Error != nil {
		return dto.Response{
			Success: false,
			Message: repoResult.Error.Error(),
		}
	}

	data, ok := repoResult.Result.(*dto.Pagination)
	if !ok || data == nil {
		return dto.Response{
			Success: false,
			Message: "invalid pagination result",
		}
	}

	urlPath := ctx.Request().URL.Path
	baseQuery := fmt.Sprintf("?limit=%d", req.Limit)

	for _, s := range req.Searchs {
		baseQuery += fmt.Sprintf("&%s.%s=%s", s.Column, s.Action, s.Query)
	}

	data.FirstPage = fmt.Sprintf("%s%s&page=1", urlPath, baseQuery)
	data.LastPage = fmt.Sprintf("%s%s&page=%d", urlPath, baseQuery, totalPages)

	if req.Page > 1 {
		data.PreviousPage = fmt.Sprintf("%s%s&page=%d", urlPath, baseQuery, req.Page-1)
	}
	if req.Page < totalPages {
		data.NextPage = fmt.Sprintf("%s%s&page=%d", urlPath, baseQuery, req.Page+1)
	}

	return dto.Response{
		Success: true,
		Data:    data,
	}
}
