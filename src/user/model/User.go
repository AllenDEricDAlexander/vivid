package model

import "time"

type User struct {
	ID              uint32    `gorm:"Column:user_id" json:"user_id"`
	Nickname        string    `gorm:"Column:user_name" json:"user_name"`
	FollowCount     int32     `gorm:"Column:follow_count" json:"follow_count"`
	FollowerCount   int32     `gorm:"Column:follower_count" json:"follower_count"`
	AvatarURL       string    `gorm:"Column:avatarURL" json:"avatarURL"`
	BackgroundImage string    `gorm:"Column:background_image" json:"background_image"`
	Signature       string    `gorm:"Column:signature" json:"signature "`
	TotalFavorite   int32     `gorm:"Column:total_favorite" json:"total_favorite"`
	WorkCount       int32     `gorm:"Column:work_count" json:"work_count"`
	FavoriteCount   int32     `gorm:"Column:favorite_count" json:"favorite_count"`
	Password        string    `gorm:"Column:password" json:"password"`
	IsDelete        int32     `gorm:"Column:is_delete" json:"is_delete"`
	CreateTime      time.Time `gorm:"Column:create_time" json:"create_time"`
	UpdateTime      time.Time `gorm:"Column:update_time" json:"update_time"`
}

func (User) TableName() string {
	return "users"
}
