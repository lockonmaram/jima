package api_entity

type GroupsCreateGroupRequest struct {
	Name       string `json:"name" validation:"required"`
	UserSerial string `json:"userSerial"`
}
type GroupsCreateGroupResponse struct {
	Serial string `json:"serial"`
	Name   string `json:"name"`
}

type GroupsAddUserToGroupRequest struct {
	GroupSerial    string `uri:"groupSerial" validation:"required"`
	UserSerial     string `uri:"userSerial" validation:"required"`
	UserAuthSerial string
}
type GroupsAddUserToGroupResponse struct {
	UserGroupSerial string `json:"userGroupSerial"`
	GroupSerial     string `json:"groupSerial"`
	UserSerial      string `json:"userSerial"`
	GroupName       string `json:"groupName"`
}

type GroupsRemoveUserFromGroupRequest struct {
	GroupSerial    string `uri:"groupSerial" validation:"required"`
	UserSerial     string `uri:"userSerial" validation:"required"`
	UserAuthSerial string
}
type GroupsRemoveUserFromGroupResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
