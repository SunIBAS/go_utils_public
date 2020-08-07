package MainTools

// 分页参数
type PageParam struct {
	// 第几页，从 0 开始
	Page int `json:"page"`
	// 每页的条目
	Count int `json:"count"`
	// 附属信息，需要在特定位置特殊解析
	Other string `json:"other"`
}
