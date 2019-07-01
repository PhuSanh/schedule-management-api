package model

type User struct {
	BaseModel
	Username 	string		`json:"username"`
	Phone 		string		`json:"phone"`
	Email 		string		`json:"email"`
	Password	string		`json:"password"`
}
