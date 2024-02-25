package proxy

import "fmt"

type IUser interface {
	GetInfo()
}

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetInfo() {
	fmt.Println("get user info")
}

type UserCachePorxy struct {
	u IUser
}

func NewUserCacheProxy(user *User) *UserCachePorxy {
	return &UserCachePorxy{u: user}
}

func (uc *UserCachePorxy) GetInfo() {
	fmt.Println("设置缓存")
	uc.u.GetInfo()
}
