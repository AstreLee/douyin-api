package common

type User struct {
	ID            uint   `gorm:"column:id"` // 主键用户ID，数据库自动生成
	Username      string // 用户名
	Password      string // 密码
	FollowCount   uint8  // 关注数量
	FollowerCount uint8  // 粉丝数量
}

// TableName 将 User 的表名设置为 `profiles`
func (User) TableName() string {
	return "user"
}
