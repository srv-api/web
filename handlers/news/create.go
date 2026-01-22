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

	userid := c.Get("UserId").(string)
	createdBy := c.Get("CreatedBy").(string)
	merchantId := c.Get("MerchantId").(string)

	title := c.FormValue("title")
	tag := c.FormValue("tag")
	body := c.FormValue("body")
	excerpt := c.FormValue("excerpt")
	status := c.FormValue("status")

	file, err := c.FormFile("file")
	if err != nil && err != http.ErrMissingFile {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	var fileName, filePath string

	if file != nil {
		uploadDir := "uploads"

		// 🔥 WAJIB
		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(c)
		}

		src, err := file.Open()
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(c)
		}
		defer src.Close()

		fileName = file.Filename
		filePath = fmt.Sprintf(
			"%s/%s_%s",
			uploadDir,
			util.GenerateRandomString(),
			file.Filename,
		)

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
		FileName:   fileName,
		FilePath:   filePath,
	}

	resp, err = h.serviceNews.Create(req)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(resp).Send(c)
}
