package constants

type Roles string

const (
	ROLES_NOAUTH Roles = ""
	ROLES_ADMIN  Roles = "administrator"
	ROLES_EMP    Roles = "employee"
	ROLES_USER   Roles = "user"
)
