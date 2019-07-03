package model

type Comment struct {
	BaseModel
	Message 	string 	`json:"message"`
	ScheduleId 	int64 	`json:"schedule_id"`
	UserId		int64 	`json:"user_id"`
}