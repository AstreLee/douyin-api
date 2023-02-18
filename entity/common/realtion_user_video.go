package common

type RelationUserVideo struct {
	ID      uint
	UserId  uint
	VideoId uint
}

func (*RelationUserVideo) TableName() string {
	return "relation_user_video"
}
