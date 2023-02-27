package model

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

type UserInfo struct {
	Email      string `json:"email" gorm:"email"`
	Name       string `json:"name" gorm:"name"`
	ScreenName string `json:"screen_name" gorm:"screen_name"`
	CreatedAt  int64  `json:"created_at" gorm:"created_at"`
	// Token      string `json:"token" gorm:"token; primary_key"`
}

type UserInfoTmp struct {
	Info  UserInfo
	Token string
}

func (u *UserInfo) GetToken() string {
	raw := []byte(strings.Join([]string{u.Email, u.ScreenName, strconv.FormatInt(u.CreatedAt, 10)}, "-"))
	h := md5.Sum(raw)
	return hex.EncodeToString(h[:])
}
