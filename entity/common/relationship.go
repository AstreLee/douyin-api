package common

type RelationShip struct {
	ID            uint `gorm:"column:id"`
	FavoriteName  string
	FavoritedName string
}
