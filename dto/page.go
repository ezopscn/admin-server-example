package dto

import "server/common"

// 数据分页响应结构体
type Page struct {
	PageNumber   uint  `json:"page_number" form:"page_number"`     // 页码
	PageSize     uint  `json:"page_size" form:"page_size"`         // 每页数据量
	Total        int64 `json:"total" form:"total"`                 // 数据量
	NoPagination bool  `json:"no_pagination" form:"no_pagination"` // 不分页，默认 false，则分页
}

// 分页数据显示格式
type PageData struct {
	Page Page        `json:"page" form:"page"` // 分页信息
	List interface{} `json:"list" form:"list"` // 数据列表
}

// 获取数据限制
func (s *Page) GetLimitAndOffset() (limit int, offset int) {
	// 如果不分页，则不偏移
	if s.NoPagination {
		return
	}

	// 如果要分页，则需要检查分页参数是否合法
	var pageSize uint
	var pageNumber uint

	// 数据量规则
	// 	1.请求数据量不能小于 1
	// 	2.请求数据量不能大于支持的最大限制
	// 这两种情况直接返回默认的每页数据量
	if s.PageSize < 1 || s.PageSize > common.MaxPageSize {
		pageSize = common.DefaultPageSize
	} else {
		pageSize = s.PageSize
	}

	// 页码规则
	pageNumber = s.PageNumber
	var total = uint(s.Total)
	if total != 0 {
		// 计算最大页码，默认取整 +1
		maxPageNumber := total/pageSize + 1
		// 如果刚好能整除，则不需要 +1
		if total%pageSize == 0 {
			maxPageNumber = total / pageSize
		}

		// 1.如果请求的页码数大于最大页码，则返回最后一页
		if s.PageNumber > maxPageNumber {
			pageNumber = maxPageNumber
		}
	}

	// 2.页码小于 1 或者每页数据则返回第一页
	if s.PageNumber < 1 {
		pageNumber = 1
	}

	// 修改请求的分页信息
	s.PageSize = pageSize
	s.PageNumber = pageNumber

	// 限制和偏移
	limit = int(pageSize)
	offset = int(pageSize * (pageNumber - 1))
	return
}
