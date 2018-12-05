package models

import "time"

type HrModel struct {
	Id  	int64 		`json:"id"`
	Name 	string 		`json:"name"` 		// 模型名称 如：cmdb
	IsLock 	bool 		`json:"is_lock"` 	// 是否锁定，锁定后不显示
	CTime  	time.Time 	`json:"c_time"`
}


