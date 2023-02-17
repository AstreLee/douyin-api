package common

type User struct {
	ID              uint   `gorm:"column:id"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	FollowCount     uint8  `json:"follow_count"`
	FollowerCount   uint8  `json:"follower_count"`
	Avatar          string `json:"avatar"`           // the relative path of head image
	BackgroundImage string `json:"background_image"` // the relative path of background image
	Signature       string `json:"signature"`        // brief introduction of user
	TotalFavorited  uint   `json:"total_favorited"`  // total favorited work
	WorkCount       uint   `json:"work_count"`       // total count work
	FavoriteCount   uint   `json:"favorite_count"`   // total favorite work
	IsFollow        bool   `json:"is_follow"`        // if follow
}

// TableName set the tableName to user
func (User) TableName() string {
	return "user"
}
