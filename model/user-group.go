package model

import "time"

type UserGroup struct {
	GroupId int64 `json:"group_id"`
	UserId 	int64 `json:"user_id"`
	AddedBy	int64 `json:"added_by"`
	CreatedAt time.Time `json:"created_at"`
}
