package entity

type Product struct {
	Id    int
	Name  string
	Price int
	Stock int
}

func (p Product) StockStatus() string {
	var status string
	if p.Stock < 3 {
		status = "Stock hampir habis"
	} else if p.Stock > 10 {
		status = "Stock terbatas"
	}

	return status
}