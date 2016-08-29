package models

type Role int

const (
	R_Guest = iota << 1
	R_User
	R_Moderator
	R_Admin
)
