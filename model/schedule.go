package model

import "time"

type Schedule struct {
	BaseModel
	UserId 			int64 		`json:"user_id"`
	GroupId 		int64 		`json:"group_id"`
	Title 			string 		`json:"title"`
	Description 	string 		`json:"description"`
	ScheduleTime 	time.Time 	`json:"schedule_time"`
	Address 		string 		`json:"address"`
	Note 			string 		`json:"note"`
	CategoryId		int64		`json:"category_id"`
}
