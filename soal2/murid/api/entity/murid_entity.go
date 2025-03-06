package entity

import "time"

type MuridEntity struct {
	ID            int      `json:"id" gorm:"primaryKey;autoIncrement"`
	Username      string    `json:"name" gorm:"type:varchar(255);not null"`
	CreateAt time.Time 		`json:"time_create" gorm:"type:timestamp"`
}