package news

import (
	"regexp"
	"strings"

	util "github.com/srv-api/util/s"
	dto "github.com/srv-api/web/dto"
)

func (s *newsService) Create(req dto.CreateNewsRequest) (dto.CreateNewsResponse, error) {
	req.ID = util.GenerateRandomString()

	// 🔥 AUTO SLUG DARI TITLE
	req.Slug = Slugify(req.Title)

	created, err := s.Repo.Create(req)
	if err != nil {
		return dto.CreateNewsResponse{}, err
	}

	return dto.CreateNewsResponse{
		ID:              created.ID,
		UserID:          created.UserID,
		MerchantID:      created.MerchantID,
		Title:           created.Title,
		Tag:             created.Tag,
		FileName:        created.FileName,
		FilePath:        created.FilePath,
		Body:            created.Body,
		Excerpt:         created.Excerpt,
		Status:          created.Status,
		CreatedBy:       created.CreatedBy,
		Slug:            created.Slug,
		MetaTitle:       created.MetaTitle,
		MetaDescription: created.MetaDescription,
	}, nil
}

func Slugify(input string) string {
	// lowercase
	slug := strings.ToLower(input)

	// ganti spasi jadi -
	slug = strings.ReplaceAll(slug, " ", "-")

	// hapus karakter aneh
	reg := regexp.MustCompile(`[^a-z0-9\-]`)
	slug = reg.ReplaceAllString(slug, "")

	// rapikan ---- jadi -
	regDash := regexp.MustCompile(`-+`)
	slug = regDash.ReplaceAllString(slug, "-")

	// trim -
	slug = strings.Trim(slug, "-")

	return slug
}
