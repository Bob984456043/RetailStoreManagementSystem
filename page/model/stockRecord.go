package model

type StockRecord struct {
	Address    string `gorm:"column:address" json:"address"`
	CreateTime int64    `gorm:"column:create_time" json:"create_time"`
	ID         int64       `gorm:"column:id;primary_key" json:"id"`
	SkuID      int64       `gorm:"column:sku_id" json:"sku_id"`
	Supplier   string `gorm:"column:supplier" json:"supplier"`
	Amount int64 	`gorm:"column:amount" json:"amount"`
	UnitPrice int64 `gorm:"column:unit_price" json:"unit_price"`
	AddUpPrice int64 `gorm:"column:add_up_price" json:"add_up_price"`
	OperationId int64 `gorm:"column:operation_id" json:"operation_id"`
}

// TableName sets the insert table name for this struct type
func (s *StockRecord) TableName() string {
	return "stock_record"
}

