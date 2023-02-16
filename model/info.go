package model

type UserInfo struct {
	Email      string `json:"email" gorm:"email"`
	Name       string `json:"name" gorm:"name"`
	ScreenName string `json:"screen_name" gorm:"screen_name"`
	Token      string `json:"token" gorm:"token,primaryKey"`
}
