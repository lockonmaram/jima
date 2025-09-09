package api_entity

type GroupCreateGroupRequest struct {
	Name       string `json:"name" validation:"required"`
	UserSerial string `json:"userSerial"`
}
type GroupCreateGroupResponse struct {
	Serial string `json:"serial"`
	Name   string `json:"name"`
}
