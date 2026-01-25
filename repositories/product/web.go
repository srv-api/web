package product

import (
	"math"

	"github.com/srv-api/merchant/entity"
	dto "github.com/srv-api/web/dto"
)

func (r *productRepository) Web(req *dto.Pagination) (RepositoryResult, int) {
	var merchants []entity.MerchantDetail
	var totalRows int64

	offset := (req.Page - 1) * req.Limit

	sort := req.Sort
	if sort == "" {
		sort = "merchant_details.created_at desc"
	}

	// ================= DATA =================
	query := r.DB.
		Model(&entity.MerchantDetail{}).
		Distinct("merchant_details.id").
		Joins("JOIN products ON products.merchant_id = merchant_details.id").
		Preload("Category").
		Preload("Product", "status = ?", 1).
		Preload("Image").
		Where("merchant_details.merchant_slug = ?", req.MerchantSlug).
		Where("products.status = ?", 1).
		Order(sort).
		Limit(req.Limit).
		Offset(offset)

	if err := query.Find(&merchants).Error; err != nil {
		return RepositoryResult{Error: err}, 0
	}

	req.Rows = merchants

	// ================= COUNT =================
	if err := r.DB.
		Model(&entity.MerchantDetail{}).
		Joins("JOIN products ON products.merchant_id = merchant_details.id").
		Where("merchant_details.merchant_slug = ?", req.MerchantSlug).
		Where("products.status = ?", 1).
		Count(&totalRows).Error; err != nil {
		return RepositoryResult{Error: err}, 0
	}

	req.TotalRows = int(totalRows)
	req.TotalPages = int(math.Ceil(float64(totalRows) / float64(req.Limit)))
	req.FromRow = offset + 1
	req.ToRow = offset + req.Limit
	if req.ToRow > int(totalRows) {
		req.ToRow = int(totalRows)
	}

	return RepositoryResult{Result: req}, req.TotalPages
}
