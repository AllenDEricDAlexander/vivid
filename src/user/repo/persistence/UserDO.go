package persistence

type UserDO struct {
	ID              uint32 `json:"user_id"`
	Nickname        string `json:"user_name"`
	FollowCount     int32  `json:"follow_count"`
	FollowerCount   int32  `json:"follower_count"`
	AvatarURL       string `json:"avatarURL"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature "`
	TotalFavorite   int32  `json:"total_favorite"`
	WorkCount       int32  `json:"work_count"`
	FavoriteCount   int32  `json:"favorite_count"`
}
