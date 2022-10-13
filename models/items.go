package models

type Items struct {
	Items_id uint `gorm:"primaryKey"`

	items_code string `gorm:"not null; unique; type: varchar(191) "`

	description string `gorm:"not null; unique; type: varchar(191) "`

	quantity int `gorm:"not null; unique; type: varchar(191) "`
}
