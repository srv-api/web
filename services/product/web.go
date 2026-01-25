package product

import (
	"fmt"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/web/dto"
)

func (s *productService) Web(ctx echo.Context, req *dto.Pagination) dto.Response {
	if req.Page < 1 {
		req.Page = 1
	}

	operationResult, totalPages := s.Repo.Web(req)

	data := operationResult.Result.(*dto.Pagination)

	urlPath := ctx.Request().URL.Path // /merchant-name
	baseQuery := fmt.Sprintf("?limit=%d", req.Limit)

	// optional search
	for _, search := range req.Searchs {
		baseQuery += fmt.Sprintf("&%s.%s=%s", search.Column, search.Action, search.Query)
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
