package models

import "time"

type HrOperate struct {
	Id  	uint64 		`json:"id"`
	Fid     uint64 		`json:"fid"`   // 功能ID
	Name 	string 		`json:"name"` 		// 名称 如：add,read,change,del
	IsLock 	bool 		`json:"is_lock"` 	// 是否锁定，锁定后不显示
	CTime  	time.Time 	`json:"c_time"`
}
