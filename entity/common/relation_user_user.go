package common

type RelationUserUser struct {
	ID          uint `gorm:"column:id"`
	FavoriteId  uint `gorm:"column:favorite_id"`
	FavoritedId uint `gorm:"column:favorited_id"`
}

func (*RelationUserUser) TableName() string {
	return "relation_user_user"
}
