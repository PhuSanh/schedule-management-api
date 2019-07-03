package model

import (
	"schedule-management-api/common"
)

type UserCategory struct {
	BaseModel
	Name 		string		`json:"name"`
	Color 		string		`json:"color"`
	Priority 	int8		`json:"priority"`
	Type		int8		`json:"type"`
	UserId		int64		`json:"user_id"`
}

const (
	DEFAULT_TYPE = 2
	DEFAULT_PRIORITY = 1
)

func (u *UserCategory) BeforeCreate() (err error) {
	if common.IsZeroOfUnderlyingType(u.Type) {
		u.Type = DEFAULT_TYPE
	}
	if common.IsZeroOfUnderlyingType(u.Priority) {
		u.Priority = DEFAULT_PRIORITY
	}
	return
}
