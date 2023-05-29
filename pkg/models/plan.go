package models

import "time"

type PlanDBModel struct {
	Id          int  `gorm:"id;primaryKey;autoIncrement"`
	Name        int  `gorm:"column:user_id;type:int;not null;uniqueIndex;not null"`
	Description int  `gorm:"column:plan_id;type:varchar(255);not null"`
	IsActive    bool `gorm:"column:is_active;type:bool;not null"`
	VoucherId   int
	CreatedAt   time.Time  `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at"`
}

func (PlanDBModel) TableName() string {
	return "plan"
}
