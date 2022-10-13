package models

import "time"

type Orders struct {
	order_id uint `gorm:"primaryKey"`

	customer_name string `gorm:"not null; unique; type: varchar(191) "`

	Ordered_at time.Time
}
