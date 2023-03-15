// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameVSupplierBill = "v_supplier_bills"

// VSupplierBill mapped from table <v_supplier_bills>
type VSupplierBill struct {
	BillID          int32     `gorm:"column:bill_id;not null;default:0" json:"bill_id"`
	SupplierID      int32     `gorm:"column:supplier_id;not null" json:"supplier_id"`
	BillStartDate   time.Time `gorm:"column:bill_start_date;not null" json:"bill_start_date"`
	BillDueDate     time.Time `gorm:"column:bill_due_date;not null" json:"bill_due_date"`
	BillNumber      string    `gorm:"column:bill_number;not null" json:"bill_number"`
	BillOrderNumber *string   `gorm:"column:bill_order_number" json:"bill_order_number"`
	BillStatus      string    `gorm:"column:bill_status;not null" json:"bill_status"`
	BillType        string    `gorm:"column:bill_type;not null" json:"bill_type"`
	SupplierName    string    `gorm:"column:supplier_name;not null" json:"supplier_name"`
	SupplierType    string    `gorm:"column:supplier_type;not null" json:"supplier_type"`
	BillTotal       float64   `gorm:"column:bill_total;not null" json:"bill_total"`
}

// TableName VSupplierBill's table name
func (*VSupplierBill) TableName() string {
	return TableNameVSupplierBill
}
