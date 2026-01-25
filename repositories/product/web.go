package product

import (
	"math"

	"github.com/srv-api/merchant/entity"
	dto "github.com/srv-api/web/dto"
	"github.com/srv-api/web/helpers"
)

func (r *productRepository) Web(req *dto.Pagination) (RepositoryResult, int) {
	var merchants []entity.MerchantDetail
	var totalRows int64

	offset := (req.Page - 1) * req.Limit

	// =========================
	// QUERY DATA
	// =========================
	query := r.DB.
		Model(&entity.MerchantDetail{}).
		Distinct("merchant_details.id").
		Joins("JOIN products ON products.merchant_detail_id = merchant_details.id").
		Preload("Category").
		Preload("Product", "status = ?", 1).
		Preload("Image").
		Where("merchant_details.merchant_slug = ?", req.MerchantSlug).
		Where("products.status = ?", 1).
		Order(req.Sort).
		Limit(req.Limit).
		Offset(offset)

	if err := query.Find(&merchants).Error; err != nil {
		return RepositoryResult{Error: err}, 0
	}

	req.Rows = merchants

	// =========================
	// COUNT (WAJIB SAMA JOIN-NYA)
	// =========================
	if err := r.DB.
		Model(&entity.MerchantDetail{}).
		Joins("JOIN products ON products.merchant_detail_id = merchant_details.id").
		Where("merchant_details.merchant_slug = ?", req.MerchantSlug).
		Where("products.status = ?", 1).
		Count(&totalRows).Error; err != nil {
		return RepositoryResult{Error: err}, 0
	}
	for i := range merchants {
		merchants[i].MerchantName = helpers.TruncateString(merchants[i].MerchantName, 47)
	}

	// =========================
	// PAGINATION INFO
	// =========================
	req.TotalRows = int(totalRows)
	req.TotalPages = int(math.Ceil(float64(totalRows) / float64(req.Limit)))

	req.FromRow = offset + 1
	req.ToRow = offset + req.Limit
	if req.ToRow > int(totalRows) {
		req.ToRow = int(totalRows)
	}

	return RepositoryResult{Result: req}, req.TotalPages
}
