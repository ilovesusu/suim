package biz

import "time"

// BaseModel 基础信息
type BaseModel struct {
	Id         int64      `gorm:"primaryKey"`
	CreateTime time.Time  `gorm:"not null;comment:创建时间"`
	UpdateTime *time.Time `gorm:"comment:更新时间"`
	DeleteTime *time.Time `gorm:"comment:删除时间"`
	Remark     string     `gorm:"size:255;comment:备注"`
}
