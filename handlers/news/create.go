package news

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	util "github.com/srv-api/util/s"
	res "github.com/srv-api/util/s/response"
	dto "github.com/srv-api/web/dto"
)

func (h *domainHandler) Create(c echo.Context) error {
	var resp dto.CreateNewsResponse

	userid, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	createdBy, ok := c.Get("CreatedBy").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	merchantId, ok := c.Get("MerchantId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	// Ambil field biasa
	title := c.FormValue("title")
	tag := c.FormValue("tag")
	body := c.FormValue("body")
	excerpt := c.FormValue("excerpt")
	status := c.FormValue("status")

	// Ambil file
	file, err := c.FormFile("file")
	if err != nil && err != http.ErrMissingFile {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	var filePath string
	if file != nil {
		// Simpan file ke folder tertentu, misal "./uploads/"
		src, err := file.Open()
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(c)
		}
		defer src.Close()

		// Bisa pakai nama unik
		filePath = fmt.Sprintf("uploads/%s_%s", util.GenerateRandomString(), file.Filename)
		dst, err := os.Create(filePath)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(c)
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(c)
		}
	}

	req := dto.CreateNewsRequest{
		UserID:     userid,
		MerchantID: merchantId,
		CreatedBy:  createdBy,
		Title:      title,
		Tag:        tag,
		Body:       body,
		Excerpt:    excerpt,
		Status:     status,
		File:       filePath, // path file yang disimpan
	}

	resp, err = h.serviceNews.Create(req)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(resp).Send(c)
}
