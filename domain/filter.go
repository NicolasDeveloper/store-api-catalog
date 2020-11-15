package domain

// Filter page representation
type Filter struct {
	Skip int `json:"skip"`
	Take int `json:"take"`
}

// NewFilter constructor
// func NewFilter(
// 	currentPage int,
// 	size int,
// 	total int,
// ) Filter {
// 	// pages := total / size

// 	return Filter{
// 		TotalItems: total,
// 	}
// }

// CurrentPage int `json:"current_page"`
// 	NextPage    int `json:"next_page"`
// 	PrevPage    int `json:"prev_page"`
