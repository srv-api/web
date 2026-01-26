package product

import (
	"math"

	"github.com/srv-api/product/entity"
	dto "github.com/srv-api/web/dto"
)

// repositories/product/web.go
func (r *productRepository) Web(req *dto.Pagination) (RepositoryResult, int) {
	var (
		merchant  entity.MerchantDetail
		products  []entity.Product
		totalRows int64
	)

	offset := (req.Page - 1) * req.Limit

	sort := req.Sort
	if sort == "" {
		sort = "products.created_at desc"
	}

	// ================= MERCHANT =================
	if err := r.DB.
		Where("merchant_slug = ?", req.MerchantSlug).
		First(&merchant).Error; err != nil {
		return RepositoryResult{Error: err}, 0
	}

	// ================= PRODUCTS =================
	if err := r.DB.
		Model(&entity.Product{}).
		Where("merchant_id = ?", merchant.ID).
		Where("status = ?", 1).
		Preload("Category").
		Preload("Image").
		Order(sort).
		Limit(req.Limit).
		Offset(offset).
		Find(&products).Error; err != nil {
		return RepositoryResult{Error: err}, 0
	}

	// ================= COUNT =================
	if err := r.DB.
		Model(&entity.Product{}).
		Where("merchant_id = ?", merchant.ID).
		Where("status = ?", 1).
		Count(&totalRows).Error; err != nil {
		return RepositoryResult{Error: err}, 0
	}

	pagination := map[string]interface{}{
		"page":        req.Page,
		"limit":       req.Limit,
		"total_rows":  totalRows,
		"total_pages": int(math.Ceil(float64(totalRows) / float64(req.Limit))),
	}

	return RepositoryResult{
		Result: map[string]interface{}{
			"merchant":   merchant,
			"products":   products,
			"pagination": pagination,
		},
	}, pagination["total_pages"].(int)
}
