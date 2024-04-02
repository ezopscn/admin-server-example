package dto

// 添加菜单表单
type AddMenuRequest struct {
	Label    string `json:"label" form:"label"`
	Icon     string `json:"icon" form:"icon"`
	Key      string `json:"key" form:"key"`
	Sort     uint   `json:"sort" form:"sort"`
	ParentId uint   `json:"parent_id" form:"parent_id"`
}

// 更新菜单表单
type UpdateMenuRequest struct {
	Id uint `json:"id" form:"id"`
	AddMenuRequest
}
