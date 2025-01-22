package enum

type Role string

const (
	RoleUser      Role = "critic"
	RoleRestorer  Role = "restorer"
	RoleModerator Role = "moderator"
	RoleAdmin     Role = "admin"
	RoleHelper    Role = "helper"
)
