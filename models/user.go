package models

type User struct {
	User_id int64
	Name    string
}

func (user *User) SetUserId(userId int64) {
	user.User_id = userId
}
