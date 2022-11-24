package common

type Pagination struct {
	Page  int
	Limit int
	Total int
}

func (p *Pagination) FullFill() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = 10
	}
}
