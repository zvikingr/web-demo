package pagination

// Pagination 为了方便从url参数里提取通用的分页信息
// 默认按sorter从小到大升序排序，reverse=true表示倒序
//
//	url: http://a.b.c/path?sort=create_time&page=3&page_size=10&reverse=true
//
//	func Handle(c *gin.Context){
//		p := new(pagination.Pagination)
//		if err := c.ShouldBindQuery(p); err != nil {
//			return
//		}
//	}
type Pagination struct {
	Sort     string `form:"sort"        json:"sort"`
	Page     int    `form:"page"        json:"page"`
	PageSize int    `form:"page_size"   json:"page_size"`
	Reverse  bool   `form:"reverse"     json:"reverse"`
}

// Valid check Pagination
func (p *Pagination) Valid() error {
	// TODO: Checksum will be added in the future
	return nil
}

// GetTerm return Calculated query conditions
func (p *Pagination) GetTerm() (int, int) {
	offset := (p.Page - 1) * p.PageSize
	limit := p.PageSize
	return offset, limit
}

// SelectAll 约定：pageSize=-1则查询所有内容
func (p *Pagination) SelectAll() bool {
	return p.PageSize == -1
}
