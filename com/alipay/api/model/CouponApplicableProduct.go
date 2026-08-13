package model

type CouponApplicableProduct struct {
	CanDelete   bool   `json:"canDelete,omitempty"`
	GmtModified string `json:"gmtModified,omitempty"`
	PriceCount  int32  `json:"priceCount,omitempty"`
	ProductId   string `json:"productId,omitempty"`
	ProductName string `json:"productName,omitempty"`
	Status      string `json:"status,omitempty"`
}
