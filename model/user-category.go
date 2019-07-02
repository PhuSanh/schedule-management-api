package model

type UserCategory struct {
	BaseModel
	Name 		string		`json:"name"`
	Color 		string		`json:"color"`
	Priority 	string		`json:"priority"`
	Type		int8		`json:"type"`
	UserId		int64		`json:"user_id"`
}