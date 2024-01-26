package models

type User struct {
	ID        int    `json:"id"`
	UUID      string `json:"uuid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	RoleId    int    `json:"roleId"`
	RoleName  string `json:"roleName"`
}
