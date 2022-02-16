package models

// User 用户信息数据结构体
type User struct {
	// ID GitHub账户id
	ID int `json:"id" bson:"id"`
	// Login GitHub账户名称
	Login string `json:"login" bson:"login"`
	// AvatarUrl 头像图片url
	AvatarUrl string `json:"avatar_url" bson:"avatar_url"`
	// Email GitHub账户邮箱
	Email string `json:"email" bson:"email"`
}
