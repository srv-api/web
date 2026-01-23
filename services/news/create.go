package news

import (
	"regexp"
	"strings"
	"unicode/utf8"

	util "github.com/srv-api/util/s"
	dto "github.com/srv-api/web/dto"
)

func (s *newsService) Create(req dto.CreateNewsRequest) (dto.CreateNewsResponse, error) {
	req.ID = util.GenerateRandomString()

	// 🔥 AUTO SEO
	req.Slug = Slugify(req.Title)
	req.MetaTitle = GenerateMetaTitle(req.Title)
	req.MetaDescription = GenerateMetaDescription(req.Excerpt, req.Body)

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

const (
	MaxMetaTitle       = 60
	MaxMetaDescription = 160
)

func GenerateMetaTitle(title string) string {
	title = strings.TrimSpace(title)

	if utf8.RuneCountInString(title) <= MaxMetaTitle {
		return title
	}

	return string([]rune(title)[:MaxMetaTitle-3]) + "..."
}

func GenerateMetaDescription(excerpt, body string) string {
	desc := strings.TrimSpace(excerpt)

	if desc == "" {
		desc = strings.TrimSpace(body)
	}

	desc = strings.ReplaceAll(desc, "\n", " ")

	if utf8.RuneCountInString(desc) <= MaxMetaDescription {
		return desc
	}

	return string([]rune(desc)[:MaxMetaDescription-3]) + "..."
}
