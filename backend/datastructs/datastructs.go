package datastructs

import "time"

type SubscriberEmail struct {
	ID        uint64 `gorm:"primaryKey"`
	Email     string
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
