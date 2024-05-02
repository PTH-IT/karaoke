package model

type User struct {
	UserID      string `gorm:"column:user_id"`
	Password    string `gorm:"column:password"`
	Email       string `gorm:"column:email"`
	Status      string `gorm:"column:status"`
	CreatedTime string `gorm:"column:created_time"`
	UpdatedTime string `gorm:"column:updated_time"`
}
type Login struct {
	UserID   string `json:"UserID"`
	Password string `json:"Password"`
}

type RegisterUser struct {
	UserID   string `json:"UserID"`
	Password string `json:"Password"`
	Email    string `json:"email"`
}
type MessageCheckUser struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
type GetUser struct {
	UserID       string `bson:"userid"`
	Password     string `bson:"password"`
	Email        string `bson:"email"`
	ActiceEmail  string `bson:"actice"`
	Status       string `bson:"status"`
	CreatedTime  string `bson:"createdtime"`
	UpdatedTime  string `bson:"updatedtime"`
	ConnectionId string `bson:"connectionid"`
}
