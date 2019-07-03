package model

type Group struct {
	BaseModel
	OwnerId	int64 	`json:"owner_id"`
	Name 	string 	`json:"name"`
}