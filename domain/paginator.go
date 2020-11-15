package domain

// Paginator page calculator
type Paginator struct {
	CurrentPageIndex int
	Total            int
	PageSize         int
	EndPagination    bool
}

//NewPaginator constructor
func NewPaginator(currentPageIndex int, pageSize int, total int) (Paginator, error) {
	endPagination := false

	if currentPageIndex > total {
		currentPageIndex = total
		endPagination = true
	}

	p := Paginator{
		CurrentPageIndex: currentPageIndex,
		PageSize:         pageSize,
		Total:            total,
		EndPagination:    endPagination,
	}

	return p, nil
}

// GetSkip get skip list
func (p *Paginator) GetSkip() int {
	return p.CurrentPageIndex
}

// GetNextPageIndex get next skip list
func (p *Paginator) GetNextPageIndex() int {
	next := p.CurrentPageIndex + p.PageSize

	if next > p.Total {
		return p.Total
	}

	return next
}

// GetPrevPageIndex get previous skip list
func (p *Paginator) GetPrevPageIndex() int {
	return p.CurrentPageIndex - p.PageSize
}

// GetTotalPageQuantity get total pages
func (p *Paginator) GetTotalPageQuantity() int {
	return p.Total / p.PageSize
}

// GetCurrentPage get current page number
func (p *Paginator) GetCurrentPage() int {
	currentPageIndex := p.CurrentPageIndex - 1

	if currentPageIndex < 0 {
		currentPageIndex = 0
	}

	return currentPageIndex / p.PageSize
}

// GetNextPage get next page number
func (p *Paginator) GetNextPage() int {
	current := p.GetCurrentPage()
	return current + 1
}

// GetPrevPage get previous page number
func (p *Paginator) GetPrevPage() int {
	return p.GetCurrentPage() - 1
}
