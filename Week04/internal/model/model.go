package model

type User struct {
	NickName string
	Gender   string
	Phone    string
	Age      int32
}

// 一些删除用户的操作可以使用这个结构体
type SimpleUser struct {
	UserID int64
}
