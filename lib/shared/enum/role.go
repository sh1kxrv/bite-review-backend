package enum

type Role string

var StaffRoles = []Role{RoleAdmin, RoleModerator}

const (
	RoleCritic    Role = "critic"
	RoleRestorer  Role = "restorer"
	RoleModerator Role = "moderator"
	RoleAdmin     Role = "admin"
	RoleHelper    Role = "helper"
)
