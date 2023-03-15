// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWholesaler = "wholesaler"

// Wholesaler mapped from table <wholesaler>
type Wholesaler struct {
	WholesalerID    int32   `gorm:"column:wholesaler_id;primaryKey;autoIncrement:true" json:"wholesaler_id"`
	ItemID          int32   `gorm:"column:item_id;not null" json:"item_id"`
	WholesalerQty   int32   `gorm:"column:wholesaler_qty;not null" json:"wholesaler_qty"`
	WholesalerPrice float64 `gorm:"column:wholesaler_price;not null" json:"wholesaler_price"`
}

// TableName Wholesaler's table name
func (*Wholesaler) TableName() string {
	return TableNameWholesaler
}
