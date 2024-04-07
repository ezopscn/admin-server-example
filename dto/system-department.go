package dto

// 添加部门表单
type AddDepartmentRequest struct {
	Name     string `json:"name" form:"name"`
	ParentId uint   `json:"parent_id" form:"parent_id"`
}

// 更新部门表单
type UpdateDepartmentRequest struct {
	Id          uint   `json:"id" form:"id"`
	ManageUsers []uint `json:"manage_users" form:"manage_users"`
	AddDepartmentRequest
}
