package common

type User struct {
	ID            uint `gorm:"column:id"`
	Username      string
	Password      string
	FollowCount   uint8
	FollowerCount uint8
}

// TableName set the tableName to user
func (User) TableName() string {
	return "user"
}
