package product

import (
	"math"

	dto "github.com/srv-api/web/dto"
)

func (r *productRepository) Web(req *dto.Pagination) (RepositoryResult, int) {
	var rows []map[string]interface{}
	var totalRows int64

	offset := (req.Page - 1) * req.Limit

	sort := req.Sort
	if sort == "" {
		sort = "merchant_details.created_at desc"
	}

	// ================= DATA =================
	err := r.DB.
		Table("merchant_details").
		Select(`
            merchant_details.id            AS merchant_id,
            merchant_details.merchant_name,
            merchant_details.merchant_slug,
            products.id                    AS product_id,
            products.product_name,
            products.price,
            products.stock,
            products.created_at            AS product_created_at
        `).
		Joins("JOIN products ON products.merchant_id = merchant_details.id").
		Where("merchant_details.merchant_slug = ?", req.MerchantSlug).
		Where("products.status = ?", 1).
		Order(sort).
		Limit(req.Limit).
		Offset(offset).
		Scan(&rows).Error

	if err != nil {
		return RepositoryResult{Error: err}, 0
	}

	req.Rows = rows

	// ================= COUNT =================
	err = r.DB.
		Table("merchant_details").
		Joins("JOIN products ON products.merchant_id = merchant_details.id").
		Where("merchant_details.merchant_slug = ?", req.MerchantSlug).
		Where("products.status = ?", 1).
		Count(&totalRows).Error

	if err != nil {
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
