// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameItemSupplier = "item_suppliers"

// ItemSupplier mapped from table <item_suppliers>
type ItemSupplier struct {
	ItemSupplierID            int32 `gorm:"column:item_supplier_id;primaryKey;autoIncrement:true" json:"item_supplier_id"`
	ItemID                    int32 `gorm:"column:item_id;not null" json:"item_id"`
	SupplierID                int32 `gorm:"column:supplier_id;not null" json:"supplier_id"`
	ItemSupplierPurchasePrice int32 `gorm:"column:item_supplier_purchase_price;not null" json:"item_supplier_purchase_price"`
	ItemSupplierSellPrice     int32 `gorm:"column:item_supplier_sell_price;not null" json:"item_supplier_sell_price"`
}

// TableName ItemSupplier's table name
func (*ItemSupplier) TableName() string {
	return TableNameItemSupplier
}
