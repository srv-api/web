package product

import (
	"math"

	merchant "github.com/srv-api/merchant/entity"
	"github.com/srv-api/product/entity"
	dto "github.com/srv-api/web/dto"
)

// repositories/product/web.go
func (r *productRepository) Web(req *dto.Pagination) (RepositoryResult, int) {
	var (
		merchant  merchant.MerchantDetail
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

	// ================= PAGINATION =================
	req.Rows = products
	req.TotalRows = int(totalRows)
	req.TotalPages = int(math.Ceil(float64(totalRows) / float64(req.Limit)))
	req.FromRow = offset + 1
	req.ToRow = offset + req.Limit
	if req.ToRow > int(totalRows) {
		req.ToRow = int(totalRows)
	}

	return RepositoryResult{
		Result: map[string]interface{}{
			"merchant":   merchant,
			"products":   products,
			"pagination": req,
		},
	}, req.TotalPages
}
