package models

import "time"

type VoucherDBModel struct {
	Id        int        `gorm:"id;primaryKey;autoIncrement"`
	Code      string     `gorm:"column:code;type:int;not null;uniqueIndex;not null"`
	ValidFor  string     `gorm:"column:valid_for;type:varchar(255);not null"`
	CreatedAt *time.Time `gorm:"column:started_at"`
}

func (VoucherDBModel) TableName() string {
	return "voucher"
}
