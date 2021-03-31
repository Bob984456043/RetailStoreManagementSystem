package model

type Sku struct {
	ImgURL        string `gorm:"column:img_url" json:"img_url"`
	Name          string      `gorm:"column:name" json:"name"`
	OriginalPrice int64     `gorm:"column:original_price" json:"original_price"`
	UpdateTime    int64   `gorm:"column:update_time" json:"update_time"`
	ID            int64       `gorm:"column:id;primary_key" json:"id"`
	CategoryID    int64       `gorm:"column:category_id" json:"category_id"`
	CreateTime    int64   `gorm:"column:create_time" json:"create_time"`
	MemberPrice   int64     `gorm:"column:member_price" json:"member_price"`
	ShopID        int64       `gorm:"column:shop_id" json:"shop_id"`
	Barcode       string      `gorm:"column:barcode" json:"barcode"`
}

// TableName sets the insert table name for this struct type
func (s *Sku) TableName() string {
	return "sku"
}

