package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Phone         string `json:"phone"`
	Password      string `json:"password"`
	NickName      string `json:"nick_name"`
	UserType      uint   `json:"user_type"`
	ValidCoin     uint   `json:"valid_coin"`
	IsVip         uint   `json:"is_vip"`
	VipExpireTime uint   `json:"vip_expire_time"`
	IsAutoRenewal uint   `json:"is_auto_renewal"`
	LastLoginTime uint   `json:"last_login_time"`
	InviteUid     uint   `json:"invite_uid"`
	InviteTime    uint   `json:"invite_time"`
	SignInDays    uint   `json:"sign_in_days"`
	LastLoginIp   string `json:"last_login_ip"`
	UserDevice    string `json:"user_device"`
	InviteCode    string `json:"invite_code"`
}