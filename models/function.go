package models

import "time"

type HrFunction struct {
	Id  	uint64 		`json:"id"`
	Mid     uint64 		`json:"mid"`   // 模型ID
	Name 	string 		`json:"name"` 		// 模型名称 如：cmdb
	IsLock 	bool 		`json:"is_lock"` 	// 是否锁定，锁定后不显示
	CTime  	time.Time 	`json:"c_time"`
}

