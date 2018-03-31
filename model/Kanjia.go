package model

import "time"

// Product 商品
type Kanjia struct {
	ID            uint       `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	DeletedAt     *time.Time `sql:"index" json:"deletedAt"`
	UserID        uint       `json:"userID"`
	UserNickName  string     `json:"userNickName"`
	UserAvatarUrl string     `json:"userAvatarUrl"`
	ProductID     uint       `json:"productID"`
	KanjiaStatus  uint       `json:"kanjiaStatus"` //0待砍价 1是砍价完成
	ValidCode     string     `json:"validCode"`    //   用于砍价活动最后验证用的编码,验证通过后砍价状态KanjiaStatus设置为1

}

type KanjiaRecord struct {
	ID            uint       `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	DeletedAt     *time.Time `sql:"index" json:"deletedAt"`
	KanjiaID      uint       `json:"kanjiaID"`
	UserID        uint       `json:"userID"`
	UserNickName  string     `json:"userNickName"`
	UserAvatarUrl string     `json:"userAvatarUrl"`
	ProductID     uint       `json:"productID"`
	KanjiaPrice   float64    `json:"kanjiaPrice"`
}
