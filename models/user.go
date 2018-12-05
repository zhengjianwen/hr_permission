package models

import "time"

type User struct {
	Id          int
	Name        string   //姓名
	UserName    string	 // 账户
	PassWord	string   // 密码
	Phone       string   // 手机号
	Position	string
	Ctime       time.Time // 创建时间
	Utime       time.Time  // 更新时间
}

