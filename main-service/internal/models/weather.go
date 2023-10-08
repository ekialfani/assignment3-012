package models

type Weather struct {
	ID int `gorm:"primaryKey" json:"-"`
	Water int `json:"water"`
	Wind int `json:"wind"`
}