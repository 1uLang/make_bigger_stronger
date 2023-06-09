/*
 * Author: lihy lihy@zhiannet.com
 * Date: 2022-11-15 11:23:31
 * LastEditors: lihy lihy@zhiannet.com
 * Note: Need note condition
 */
package tools

type Page struct {
	curPage   int
	size      int
	totalSzie int
	totalPage int
}

func NewPage(curPage int, size int, total int) *Page {
	page := &Page{
		curPage:   curPage,
		size:      size,
		totalSzie: total,
	}

	if page.size > 100 {
		page.size = 100
	}

	// page.totalPage = int(math.Ceil(float64(total) / float64(size)))
	// if page.curPage > page.totalPage {
	// 	page.curPage = page.totalPage
	// }

	return page
}

func (p *Page) Page() int {
	return p.curPage + 1
}

func (p *Page) Offset() int {
	if p.curPage == 0 {
		p.curPage = 1
	}
	return (p.curPage - 1) * p.size
}

func (p *Page) Size() int {
	return p.size
}

/*
 * Author: git config user.email
 * file name: File name
 * Data: Do not edit
 * LastEditor:
 * LastData:
 * Describe:
 */
func (p *Page) TotalSize() int {
	return p.totalSzie
}

func (p *Page) TotalPage() int {
	return p.totalPage
}
