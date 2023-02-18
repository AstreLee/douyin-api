package common

type Video struct {
	ID            uint   `json:"id" gorm:"column: id"`
	Author        User   `json:"author" gorm:"foreignKey: ID"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount uint   `json:"favorite_count"`
	CommentCount  uint   `json:"comment_count"`
	Title         string `json:"title"`
	IsFavorite    bool   `json:"is_favorite"`
}

func (*Video) TableName() string {
	return "video"
}
