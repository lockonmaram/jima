package api_entity

type Group struct {
	GroupSerial string `json:"groupSerial"`
	Name        string `json:"name"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
type GroupMember struct {
	UserGroupSerial string `json:"userGroupSerial"`
	UserSerial      string `json:"userSerial"`
	UserName        string `json:"userName"`
	Role            string `json:"role"`
	MemberSince     string `json:"memberSince"`
}

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

type GroupsGetGroupBySerialRequest struct {
	GroupSerial    string `uri:"groupSerial" validation:"required"`
	UserAuthSerial string
}
type GroupsGetGroupBySerialResponse struct {
	Group
}

type GroupsGetGroupsByUserSerialRequest struct {
	UserSerial string `uri:"userSerial" validation:"required"`
}
type GroupsGetGroupsByUserSerialResponse struct {
	Groups []Group `json:"groups"`
}

type GroupsGetGroupMembersRequest struct {
	GroupSerial    string `uri:"groupSerial" validation:"required"`
	UserAuthSerial string
}
type GroupsGetGroupMembersResponse struct {
	GroupMembers []GroupMember `json:"groupMembers"`
}

type GroupsUpdateGroupRequest struct {
	GroupSerial    string `uri:"groupSerial" validation:"required"`
	Name           string `json:"name"`
	UserAuthSerial string
}
type GroupsUpdateGroupResponse struct {
	Group
}

type GroupsUpdateGroupMemberRoleRequest struct {
	GroupSerial    string `uri:"groupSerial" validation:"required"`
	UserSerial     string `uri:"userSerial" validation:"required"`
	Role           string `json:"role"`
	UserAuthSerial string
}
type GroupsUpdateGroupMemberRoleResponse struct {
	GroupMember
}
