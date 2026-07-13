package link

import (
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

func NewLink(url string, hash string) *Link {
	return &Link{
		Url:  url,
		Hash: hash,
	}
}



