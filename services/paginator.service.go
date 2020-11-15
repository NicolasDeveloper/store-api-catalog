package services

// CalculateSkip calculate skip list
func CalculateSkip(currentPageIndex int, pageSize int) int {
	return currentPageIndex + pageSize
}

// CalculateNextPageIndex calculate next skip list
func CalculateNextPageIndex(currentPageIndex int, pageSize int) int {
	return currentPageIndex + (pageSize * 2)
}

// CalculatePrevPageIndex calculate previous skip list
func CalculatePrevPageIndex(currentPageIndex int, pageSize int) int {
	return currentPageIndex - pageSize
}

// CalculateTotalPageQuantity calculate total pages
func CalculateTotalPageQuantity(pageSize int, total int) int {
	return total / pageSize
}

// CalculateCurrentPage calculate current page number
func CalculateCurrentPage(currentPageIndex int, pageSize int, total int) int {
	return ((currentPageIndex - 1) / pageSize) + 1
}

// CalculateNextPage calculate next page number
func CalculateNextPage(currentPageIndex int, pageSize int, total int) int {
	return CalculateCurrentPage(currentPageIndex, pageSize, total) + 1
}

// CalculatePrevPage calculate previous page number
func CalculatePrevPage(currentPageIndex int, pageSize int, total int) int {
	return CalculateCurrentPage(currentPageIndex, pageSize, total) - 1
}
